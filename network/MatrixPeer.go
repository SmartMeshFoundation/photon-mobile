package network

import (
	"fmt"

	"github.com/SmartMeshFoundation/Photon/utils"

	"github.com/SmartMeshFoundation/Photon/log"

	"github.com/SmartMeshFoundation/Photon/network/gomatrix"
	"github.com/ethereum/go-ethereum/common"
)

type peerStatus int

const (
	peerStatusUnkown = iota
	peerStatusOffline
	peerStatusOnline
)

func (s peerStatus) String() string {
	switch s {
	case peerStatusUnkown:
		return UNAVAILABLE
	case peerStatusOffline:
		return OFFLINE
	case peerStatusOnline:
		return ONLINE
	}
	return "error status"
}

//MatrixPeer is the  photon node on matrix server
type MatrixPeer struct {
	address common.Address //需要通信的对象
	//address 对应的所有可能的 User
	candidateUsers       map[string]*gomatrix.UserInfo
	candidateUsersStatus map[string]peerStatus
	//确定对方在线的那个聊天室
	defaultMessageRoomID string
	rooms                map[string]bool //roomID exists?
	status               peerStatus
	deviceType           string
	receiveMessage       chan struct{}
	channelCount         int // 我与此节点总共有多少条通道
}

//NewMatrixPeer create matrix user
func NewMatrixPeer(address common.Address) *MatrixPeer {
	u := &MatrixPeer{
		address:              address,
		rooms:                make(map[string]bool),
		candidateUsers:       make(map[string]*gomatrix.UserInfo),
		candidateUsersStatus: make(map[string]peerStatus),
		channelCount:         1,
	}
	return u
}

func (peer *MatrixPeer) isValidUserID(userID string) bool {
	for _, u := range peer.candidateUsers {
		if u.UserID == userID {
			return true
		}
	}
	return false
}

/*
make sure that `users` are already joined in the default message room
*/
func (peer *MatrixPeer) updateUsers(users []*gomatrix.UserInfo) {
	for _, lu := range users {
		peer.candidateUsers[lu.UserID] = lu
	}
	return
}

func (peer *MatrixPeer) addRoom(roomID string) {
	if peer.rooms[roomID] {
		log.Warn(fmt.Sprintf("roomID %s already exists for %s", roomID, utils.APex(peer.address)))
	}
	peer.rooms[roomID] = true
}

/*
1. if one of the `candidateUsers` is online, status is online ,otherwise
2. if one of the `candidateUsers` is offline,status is offline
3. then if all of the `candidateUsers` is UNAVAILABLE,status is unkown
*/
func (peer *MatrixPeer) setStatus(userID string, presence string) bool {
	peer.candidateUsers[userID] = &gomatrix.UserInfo{
		DisplayName: "",
		AvatarURL:   "",
		UserID:      userID,
	}
	var status peerStatus
	switch presence {
	case ONLINE:
		status = peerStatusOnline
	case OFFLINE:
		status = peerStatusOffline
	case UNAVAILABLE:
		status = peerStatusUnkown
	}
	user := peer.candidateUsers[userID]
	if user == nil {
		log.Error(fmt.Sprintf("peer %s ,userid %s set status %s ,but i don't kown this userid. MatrixPeer=%s",
			utils.APex2(peer.address), userID, status, utils.StringInterface(peer, 7)))
		return false
	}
	peer.candidateUsersStatus[userID] = status
	for _, s := range peer.candidateUsersStatus {
		if s > status {
			return false
		}
	}
	peer.status = status
	return true
}

//如果小于等于0,说明已经没有任何channel了,这个节点可以移除.
func (peer *MatrixPeer) decreaseChannelCount() bool {
	peer.channelCount--
	return peer.channelCount <= 0
}

func (peer *MatrixPeer) increaseChannelCount() {
	peer.channelCount++
}
