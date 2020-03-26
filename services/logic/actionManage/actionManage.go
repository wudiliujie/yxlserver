package actionManage

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/network"
	"yxlserver/services/I"
	"yxlserver/services/logic/playerManage"
	"yxlserver/services/util"
)

var pckHandleAgentMap = make(map[uint16]func(agent gate.Agent, pckMsg network.IMessage))
var pckHandlePlayerMap = make(map[uint16]func(player I.IPlayer, pckMsg network.IMessage))

func RegisterAgentPacket(msgId uint16, f func(agent gate.Agent, pckMsg network.IMessage)) {
	pckHandleAgentMap[msgId] = f
}
func RegisterPlayerPacket(msgId uint16, f func(player I.IPlayer, pckMsg network.IMessage)) {
	pckHandlePlayerMap[msgId] = f
}
func ActionAgentHandle(agent gate.Agent, msgId uint16, pck network.IMessage) {
	//如果登录成功
	userId := agent.UserData().(int64)
	if userId > 0 {
		player := playerManage.GetPlayerById(userId)
		if util.IsNil(player) == false {
			f, ok := pckHandlePlayerMap[msgId]
			if ok {
				f(player, pck)
			} else {
				log.Error("ActionAgentHandle msg:%v 没有注册 %v", msgId, agent.RemoteAddr().String())
			}
		} else {
			log.Error("玩家不存在：%v", userId)
		}

	} else {
		f, ok := pckHandleAgentMap[msgId]
		if ok {
			f(agent, pck)
		} else {
			log.Error("ActionAgentHandle msg:%v 没有注册 %v", msgId, agent.RemoteAddr().String())
		}
	}
}
