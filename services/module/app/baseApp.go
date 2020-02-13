package app

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"yxlserver/services/I"
	"yxlserver/services/I/eventcode"
	"yxlserver/services/logic/actionManage"
)

type BaseAppModule struct {
	*module.Skeleton
	self     I.IAppModule
	day      int
	hour     int
	minute   int
	agentMap map[int32]gate.Agent
}

func (m *BaseAppModule) OnInit() {
	m.Skeleton = module.NewSkeleton()
	m.agentMap = make(map[int32]gate.Agent)
	module.RegisterChanRpcEvent(eventcode.Net_AgentInit, m.ChanRPCServer, m.self.OnAgentInit)
	module.RegisterChanRpcEvent(eventcode.Net_AgentDestroy, m.ChanRPCServer, m.self.OnAgentDestroy)
	module.RegisterChanRpcEvent(eventcode.Net_ReceiveMsg, m.ChanRPCServer, m.self.OnReceiveMsg)
	module.RegisterChanRpcEvent(module.OnInitComplete, m.ChanRPCServer, m.self.OnInitComplete)

}
func (m *BaseAppModule) OnDestroy() {

}
func (m *BaseAppModule) OnAgentInit(args []interface{}) {
	agent := args[0].(gate.Agent)
	m.agentMap[agent.GetId()] = agent
}
func (m *BaseAppModule) OnAgentDestroy(args []interface{}) {
	agent := args[0].(gate.Agent)
	delete(m.agentMap, agent.GetId())
}

func (m *BaseAppModule) OnReceiveMsg(args []interface{}) {
	//重置心跳
	agent := args[0].(gate.Agent)
	log.Debug("BaseAppModule OnReceiveMsg")
	id := args[1].(uint16)
	agent.ResetHeart()
	actionManage.ActionAgentHandle(agent, id, args[2])

}
func (m *BaseAppModule) OnInitComplete(args []interface{}) {
	log.Debug("初始化完成")
}
