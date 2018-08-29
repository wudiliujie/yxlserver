package internal

import (
	"common/msg"
	"fmt"
	"gameserver/app"
	"gameserver/conf"
	"gameserver/events"
	"leaf/gate"
	"leaf/log"
)

type Module struct {
	*gate.Gate
}

var appChanRPC = app.ChanRPC

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:         conf.Server.MaxConnNum,
		PendingWriteNum:    conf.PendingWriteNum,
		MaxMsgLen:          conf.MaxMsgLen,
		WSAddr:             conf.Server.WSAddr,
		HTTPTimeout:        conf.HTTPTimeout,
		CertFile:           conf.Server.CertFile,
		KeyFile:            conf.Server.KeyFile,
		TCPAddr:            conf.Server.TCPAddr,
		LenMsgLen:          conf.LenMsgLen,
		LittleEndian:       conf.LittleEndian,
		Processor:          msg.Processor,
		GoLen:              conf.AgentGoLen,
		TimerDispatcherLen: conf.AgentTimerDispatcherLen,
		AsynCallLen:        conf.AgentAsynCallLen,
		ChanRPCLen:         conf.AgentChanRPCLen,
		OnAgentInit:        onAgentInit,
		OnAgentDestroy:     onAgentDestroy,
		OnReceiveMsg:       onReceiveMsg,
	}
	m.Gate.AgentChanRPC = app.ChanRPC
}
func onAgentInit(agent gate.Agent) {
	fmt.Print("客户端连接:%v", agent.RemoteAddr())
	appChanRPC.Go(events.Net_AgentInit, agent)

}

func onAgentDestroy(agent gate.Agent) {
	fmt.Print("客户端断开连接:%v", agent.RemoteAddr())
	appChanRPC.Go(events.Net_AgentDestroy, agent)
}
func (m *Module) OnDestroy() {
	log.Debug("销毁Gate")
}
func onReceiveMsg(agent gate.Agent, msgid int32, pck interface{}) {
	appChanRPC.Go(events.Net_ReceiveMsg, agent, msgid, pck)
}
