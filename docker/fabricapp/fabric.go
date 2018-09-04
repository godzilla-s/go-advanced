package main

import "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

type FabricSetup struct {
	ConfigFile    string                           //sdk配置文件所在路径
	ChannelID     string                           //应用通道名称
	ChannelConfig string                           //应用通道交易配置文件所在路径
	OrgAdmin      string                           // 组织管理员名称
	OrgName       string                           //组织名称
	Initialized   bool                             //是否初始化
	Admin         resmgmtclient.ResourceMgmtClient //fabric环境中资源管理者
	SDK           *fabsdk.FabricSDK                //SDK实例
}

func (f *FabricSetup) Init() {
	fab.New(nil)
}

func main() {
	f := new(FabricSetup)

	f.Init()
}
