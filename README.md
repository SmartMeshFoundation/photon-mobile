# photon-mobile
[![Build Status](https://travis-ci.org/SmartMeshFoundation/photon-mobile.svg?branch=master)](https://github.com/SmartMeshFoundation/photon-mobile)
[![Go Report Card](https://goreportcard.com/badge/github.com/SmartMeshFoundation/photon-mobile)](https://github.com/SmartMeshFoundation/photon-mobile)



![](https://github.com/dognie/Photon/blob/master/docs/photon.png?raw=true)

 [Photon documentation](https://PhotonNetwork.readthedocs.io/en/latest/)

Photon-mobile is a customized version of photon for mobile mobile platforms, which is mainly used for mobile node payment. The biggest difference with photon is that it is not considered as an intermediate node for payment.
 
## Project Status
  This project is still very much a work in progress. It can be used for testing, but it should not be used for real funds.  
## Build
 

## mobile support
Photon can works on Android and iOS using mobile's API.  it needs [go mobile](https://github.com/golang/mobile) to build mobile library.
### build Android mobile library
```bash
cd mobile
./build_Android.sh 
```
then you can integrate `mobile.aar` into your project.
### build iOS mobile framework
```bash
./build_iOS.sh
```
then you can integrate `Mobile.framework` into your project.
## Requirements
Latest version of SMC

We need go's plugin module.
