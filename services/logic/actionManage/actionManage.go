package actionManage

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
)

var pckHandleAgentMap = make(map[uint16]func(agent gate.Agent, pckMsg interface{}))

func RegisterAgentPacket(msgId uint16, f func(agent gate.Agent, pckMsg interface{})) {
	pckHandleAgentMap[msgId] = f
}

func ActionAgentHandle(agent gate.Agent, msgId uint16, pck interface{}) {

	f, ok := pckHandleAgentMap[msgId]
	if ok {
		f(agent, pck)
	} else {
		log.Error("ActionAgentHandle msg:%v 没有注册 %v", msgId, agent.RemoteAddr().String())
	}
}
