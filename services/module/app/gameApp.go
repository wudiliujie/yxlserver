package app

import (
	"github.com/wudiliujie/common/module"
	"yxlserver/services/I/eventcode"
	"yxlserver/services/logic/areaManage"
	"yxlserver/services/logic/playerManage"
)

type GameAppModule struct {
	BaseAppModule
}

func CreateGameAppModule() *GameAppModule {
	ret := &GameAppModule{}
	ret.self = ret
	return ret
}
func (m *GameAppModule) OnInit() {
	m.BaseAppModule.OnInit()
	m.Skeleton.Engine = m
	module.RegisterChanRpcEvent(eventcode.DB_LoadRoleData_Callback, m.ChanRPCServer, playerManage.LoadUserCallBack)
	initCommand(m.Skeleton)
	areaManage.Init()
}

func (m *GameAppModule) Update(diff int64) {
	playerManage.Update()
}
