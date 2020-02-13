package app

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
	"yxlserver/services/logic/serverManage"
)

type GateAppModule struct {
	BaseAppModule
}

func CreateGateAppModule() *GateAppModule {
	ret := &GateAppModule{}
	ret.self = ret
	return ret
}
func (m *GateAppModule) OnAgentInit(args []interface{}) {
	m.BaseAppModule.OnAgentInit(args)
	agent := args[0].(gate.Agent)
	if agent.GetTag() > 0 { //是内部链接
		serverManage.OnServerConnect(agent)
	}
}

func (m *GateAppModule) OnInitComplete(args []interface{}) {
	log.Debug("GateAppModule 初始化完成")
	serverManage.ConnectCenterServer()
}
