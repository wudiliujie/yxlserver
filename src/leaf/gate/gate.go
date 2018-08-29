package gate

import (
	"leaf/chanrpc"
	"leaf/log"
	"leaf/module"
	"leaf/network"
	"net"
	"reflect"
	"sync/atomic"
	"time"
)

var sessionId int32 = 0

type Gate struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network.Processor
	AgentChanRPC    *chanrpc.Server

	// websocket
	WSAddr      string
	HTTPTimeout time.Duration
	CertFile    string
	KeyFile     string

	// tcp
	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool

	// agent
	GoLen              int
	TimerDispatcherLen int
	AsynCallLen        int
	ChanRPCLen         int
	OnAgentInit        func(Agent)
	OnAgentDestroy     func(Agent)
	OnReceiveMsg       func(agent Agent, msgId int32, pck interface{})
}

func (gate *Gate) Run(closeSig chan bool) {
	newAgent := func(conn network.Conn) network.Agent {
		a := &agent{conn: conn, gate: gate}
		a.id = atomic.AddInt32(&sessionId, 1)
		if gate.OnAgentInit != nil {
			gate.OnAgentInit(a)
		}
		return a
	}

	var wsServer *network.WSServer
	if gate.WSAddr != "" {
		wsServer = new(network.WSServer)
		wsServer.Addr = gate.WSAddr
		wsServer.MaxConnNum = gate.MaxConnNum
		wsServer.PendingWriteNum = gate.PendingWriteNum
		wsServer.MaxMsgLen = gate.MaxMsgLen
		wsServer.HTTPTimeout = gate.HTTPTimeout
		wsServer.CertFile = gate.CertFile
		wsServer.KeyFile = gate.KeyFile
		wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
			return newAgent(conn)
		}
	}

	var tcpServer *network.TCPServer
	if gate.TCPAddr != "" {
		tcpServer = new(network.TCPServer)
		tcpServer.Addr = gate.TCPAddr
		tcpServer.MaxConnNum = gate.MaxConnNum
		tcpServer.PendingWriteNum = gate.PendingWriteNum
		tcpServer.LenMsgLen = gate.LenMsgLen
		tcpServer.MaxMsgLen = gate.MaxMsgLen
		tcpServer.LittleEndian = gate.LittleEndian
		tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
			return newAgent(conn)
		}
	}

	if wsServer != nil {
		wsServer.Start()
	}
	if tcpServer != nil {
		tcpServer.Start()
	}
	<-closeSig
	if wsServer != nil {
		wsServer.Close()
	}
	if tcpServer != nil {
		tcpServer.Close()
	}
}

func (gate *Gate) OnDestroy() {}

type agent struct {
	conn     network.Conn
	gate     *Gate
	userData interface{}
	module   module.Module
	id       int32
}

func (a *agent) Run() {
	closeSig := make(chan bool, 1)
	defer func() {
		if r := recover(); r != nil {
			log.Recover(r)
		}

		closeSig <- true
	}()

	handleMsgData := func(args []interface{}) error {
		if a.gate.Processor != nil {
			data := args[0].([]byte)
			msg, err := a.gate.Processor.Unmarshal(data)

			if err != nil {
				return err
			}
			msgId := a.gate.Processor.GetMsgId(msg)
			if a.gate.OnReceiveMsg != nil {
				a.gate.OnReceiveMsg(a, msgId, msg)
			}
		}
		return nil
	}
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}
		//等待切换module

		err = handleMsgData([]interface{}{data})
		if err != nil {
			log.Debug("handle message: %v", err)
			break
		}
	}
}

func (a *agent) OnClose() {
	if a.gate.OnAgentDestroy != nil {
		//log.Debug("close agent")
		a.gate.OnAgentDestroy(a)
	}
}

func (a *agent) WriteMsg(msg interface{}) {
	if a.gate.Processor != nil {
		data, err := a.gate.Processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agent) Close() {
	a.conn.Close()
}

func (a *agent) Destroy() {
	a.conn.Destroy()
}

func (a *agent) UserData() interface{} {
	return a.userData
}

func (a *agent) SetUserData(data interface{}) {
	a.userData = data
}
func (a *agent) Gate() *Gate {
	return a.gate
}
func (a *agent) GetId() int32 {
	return a.id
}
