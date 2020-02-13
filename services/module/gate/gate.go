package gate

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"yxlserver/services/I/eventcode"
	"yxlserver/services/conf"
	"yxlserver/services/gconf"
	"yxlserver/services/logic/serverManage"
	"yxlserver/services/msg"
)

var Module = new(GateModule)
var netEvent = new(gate.NetEvent)

func init() {
	gconf.Processor = msg.Processor
	netEvent.OnAgentInit = onAgentInit
	netEvent.OnAgentDestroy = onAgentDestroy
	netEvent.OnReceiveMsg = onReceiveMsg
	netEvent.Processor = gconf.Processor
	serverManage.SetNetEvent(netEvent)
}

type GateModule struct {
	*gate.Gate
}

func (m *GateModule) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		WSAddr:          conf.Server.WSAddr,
		PendingWriteNum: gconf.PendingWriteNum,
		MaxMsgLen:       gconf.MaxMsgLen,
		HTTPTimeout:     gconf.HTTPTimeout,
		//CertFile:      conf.Server.CertFile,
		//KeyFile:       conf.Server.KeyFile,
		//TCPAddr:       conf.Server.TCPAddr,
		LenMsgLen:    gconf.LenMsgLen,
		LittleEndian: gconf.LittleEndian,

		GoLen:              gconf.AgentGoLen,
		TimerDispatcherLen: gconf.AgentTimerDispatcherLen,
		AsynCallLen:        gconf.AgentAsyncCallLen,
		ChanRPCLen:         gconf.AgentChanRPCLen,
		NetEvent:           netEvent,
	}

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

func onReceiveMsg(agent gate.Agent, data []byte) {
	if gconf.Processor != nil {
		if len(data) < 4 {
			log.Debug("长度不够")
			return
		}
		//根据agent类型不同，解析不同的头部

		pck, err := gconf.Processor.Unmarshal(data[2:])
		if err != nil {
			log.Error("onReceiveMsg %v", err)
		}
		log.Debug("*******************%v", pck)
		module.OnChanRpcEvent(eventcode.Net_ReceiveMsg, agent, pck.GetId(), pck)
	}

}
