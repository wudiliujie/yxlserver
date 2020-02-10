package gate

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"yxlserver/services/I/eventcode"
	"yxlserver/services/gconf"
)

var Module = new(GateModule)

type GateModule struct {
	*gate.Gate
}

func (m *GateModule) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      gconf.MaxConnNum,
		PendingWriteNum: gconf.PendingWriteNum,
		MaxMsgLen:       gconf.MaxMsgLen,
		WSAddr:          gconf.WSAddr,
		HTTPTimeout:     gconf.HTTPTimeout,
		//CertFile:        conf.Server.CertFile,
		//KeyFile:         conf.Server.KeyFile,
		//TCPAddr:            conf.Server.TCPAddr,
		LenMsgLen:    gconf.LenMsgLen,
		LittleEndian: gconf.LittleEndian,
		Processor:    gconf.Processor,

		GoLen:              gconf.AgentGoLen,
		TimerDispatcherLen: gconf.AgentTimerDispatcherLen,
		AsynCallLen:        gconf.AgentAsynCallLen,
		ChanRPCLen:         gconf.AgentChanRPCLen,
		OnAgentInit:        onAgentInit,
		OnAgentDestroy:     onAgentDestroy,
		OnReceiveMsg:       onReceiveMsg,
	}
	gate.ThisGate = m.Gate

}
func onAgentInit(agent gate.Agent) {
	log.Debug("客户端连接:%v", agent.RemoteAddr())
	//怎么通知
	module.OnChanRpcEvent(eventcode.Net_AgentInit, agent)
}

func onAgentDestroy(agent gate.Agent) {
	log.Debug("客户端断开连接:%v", agent.RemoteAddr())
	module.OnChanRpcEvent(eventcode.Net_AgentDestroy, agent)
}
func (m *GateModule) OnDestroy() {
	log.Debug("销毁Gate")
}
func (m *GateModule) Debug() {
	log.Debug("gate")
}

func onReceiveMsg(agent gate.Agent, msgid uint16, pck interface{}) {
	//这里过滤下消息
	module.OnChanRpcEvent(eventcode.Net_ReceiveMsg, agent, int32(msgid), pck)
}
