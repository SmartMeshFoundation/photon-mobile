// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 文件空间管理
package gfilespace

import (
	"gitee.com/johng/gkvdb/gkvdb/gbtree"
	"sync"
)

// 文件空间管理结构体
type Space struct {
	mu      sync.RWMutex          // 并发操作锁
	blocks  *gbtree.BTree         // 所有的空间块构建的B+树
	sizetr  *gbtree.BTree         // 空间块大小构建的B+树
	sizemap map[int]*gbtree.BTree // 按照空间块大小构建的索引哈希表，便于检索，每个表项是一个B+树
}

// 文件空闲块
type Block struct {
	index int // 文件偏移量
	size  int // 区块大小(byte)
}

// 用于B+树的接口具体实现定义
func (block *Block) Less(item gbtree.Item) bool {
	if block.index < item.(*Block).index {
		return true
	}
	return false
}

// 创建一个空间管理器
func New() *Space {
	return &Space{
		blocks:  gbtree.New(10),
		sizetr:  gbtree.New(5),
		sizemap: make(map[int]*gbtree.BTree),
	}
}

// 添加空闲空间到管理器
func (space *Space) addBlock(index int, size int) {
	block := &Block{index, size}

	// 插入进全局树
	space.blocks.ReplaceOrInsert(block)

	// 插入进入索引表
	space.insertIntoSizeMap(block)

	// 对插入的数据进行合并检测
	space.checkMerge(block)
}

// 获取指定block的前一项block
func (space *Space) getPrevBlock(block *Block) *Block {
	var pblock *Block = nil
	space.blocks.DescendLessOrEqual(block, func(item gbtree.Item) bool {
		if item.(*Block).index != block.index {
			pblock = item.(*Block)
			return false
		}
		return true
	})
	return pblock
}

// 获取指定block的后一项block
func (space *Space) getNextBlock(block *Block) *Block {
	var nblock *Block = nil
	space.blocks.AscendGreaterOrEqual(block, func(item gbtree.Item) bool {
		if item.(*Block).index != block.index {
			nblock = item.(*Block)
			return false
		}
		return true
	})
	return nblock
}

// 获取指定block的前一项block size
func (space *Space) getPrevBlockSize(size int) int {
	psize := 0
	space.sizetr.DescendLessOrEqual(gbtree.Int(size), func(item gbtree.Item) bool {
		if int(item.(gbtree.Int)) != size {
			psize = int(item.(gbtree.Int))
			return false
		}
		return true
	})
	return psize
}

// 获取指定block的后一项block size
func (space *Space) getNextBlockSize(size int) int {
	nsize := 0
	space.sizetr.AscendGreaterOrEqual(gbtree.Int(size), func(item gbtree.Item) bool {
		if int(item.(gbtree.Int)) != size {
			nsize = int(item.(gbtree.Int))
			return false
		}
		return true
	})
	return nsize
}

// 内部按照索引检查合并
func (space *Space) checkMerge(block *Block) {
	// 首先检查插入空间块的前一项往后是否可以合并，如果当前合并失败后，才会判断当前插入项和后续的空间块合并
	if b := space.checkMergeOfTwoBlock(space.getPrevBlock(block), block); b.index == block.index {
		// 其次检查插入空间块的当前项往后是否可以合并
		space.checkMergeOfTwoBlock(block, space.getNextBlock(block))
	}
}

// 连续检测两个空间块的合并，返回最后一个无法合并的空间块指针
func (space *Space) checkMergeOfTwoBlock(pblock, block *Block) *Block {
	if pblock == nil {
		return block
	}
	if block == nil {
		return pblock
	}
	for {
		if pblock.index+int(pblock.size) >= block.index {
			space.removeBlock(block)
			// 判断是否需要更新大小
			if pblock.index+int(pblock.size) < block.index+int(block.size) {
				space.removeFromSizeMap(pblock)
				pblock.size = block.index + block.size - pblock.index
				space.insertIntoSizeMap(pblock)
			}
			block = space.getNextBlock(pblock)
			if block == nil {
				return pblock
			}
		} else {
			break
		}
	}
	return block
}

// 插入空间块到索引表
func (space *Space) insertIntoSizeMap(block *Block) {
	tree, ok := space.sizemap[block.size]
	if !ok {
		tree = gbtree.New(10)
		space.sizemap[block.size] = tree
	}
	tree.ReplaceOrInsert(block)

	// 插入空间块大小记录表
	space.sizetr.ReplaceOrInsert(gbtree.Int(block.size))
}

// 删除一项
func (space *Space) removeBlock(block *Block) {
	space.blocks.Delete(block)
	space.removeFromSizeMap(block)
}

// 从索引表中删除对应的空间块
func (space *Space) removeFromSizeMap(block *Block) {
	if tree, ok := space.sizemap[block.size]; ok {
		tree.Delete(block)
		// 数据数据为空，那么删除该项哈希记录
		if tree.Len() == 0 {
			delete(space.sizemap, block.size)
			space.sizetr.Delete(gbtree.Int(block.size))
		}
	}
}

// 获得碎片偏移量
func (block *Block) Index() int {
	return block.index
}

// 获得碎片大小
func (block *Block) Size() int {
	return block.size
}
