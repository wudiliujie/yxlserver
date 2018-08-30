package internal

import (
	"gameserver/common/agentManage"
	"gameserver/common/playerManage"
	"gameserver/events"
	"leaf/gate"
	"leaf/util"
)

func init() {
	ChanRPC.Register(events.Net_AgentInit, OnAgentInit)
	ChanRPC.Register(events.Net_AgentDestroy, OnAgentDestroy)
	ChanRPC.Register(events.Net_ReceiveMsg, OnReceiveMsg)
}
func OnAgentInit(args []interface{}) {
	agent := args[0].(gate.Agent)
	agentManage.AddAgent(agent)
}
func OnAgentDestroy(args []interface{}) {
	agent := args[0].(gate.Agent)
	agentManage.RemoveAgent(agent)
}
func OnReceiveMsg(args []interface{}) {
	//log.Debug("onrece:%v", args)
	agent := args[0].(gate.Agent)
	msgId := args[1].(int32)
	pck := args[2]
	var player *playerManage.Player
	if agent.UserData() == nil {
		userId := util.RandInterval(1, 100)
		player = playerManage.NewPlayer(userId)
		agent.SetUserData(player)
	} else {
		player = agent.UserData().(*playerManage.Player)
	}
	player.AddNetMsg(msgId, pck)
}
