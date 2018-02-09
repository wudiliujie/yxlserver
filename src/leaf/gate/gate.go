package gate

import (
	"leaf/chanrpc"
	"leaf/log"
	"leaf/network"
	"net"
	"reflect"
	"time"
	"leaf/module"
)

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
	OnAgentInit 	   func(Agent)
	OnAgentDestroy 	   func(Agent)
}

func (gate *Gate) Run(closeSig chan bool) {
	newAgent := func(conn network.Conn) network.Agent {
		a := &agent{conn: conn, gate: gate}

	/*	if gate.ChanRPCLen > 0 {
			skeleton := &module.Skeleton{
				GoLen:              gate.GoLen,
				TimerDispatcherLen: gate.TimerDispatcherLen,
				AsynCallLen:        gate.AsynCallLen,
				ChanRPCServer:      chanrpc.NewServer(gate.ChanRPCLen),
			}
			skeleton.Init()

			a.skeleton = skeleton
			a.chanRPC = skeleton.ChanRPCServer
		}*/
		if gate.AgentChanRPC != nil {
			a.chanRPC = gate.AgentChanRPC
			gate.AgentChanRPC.Go("NewAgent", a)
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
	//skeleton *module.Skeleton
	chanRPC  *chanrpc.Server
	gate     *Gate
	userData interface{}
	module  module.Module
}

func (a *agent) Run() {
	closeSig := make(chan bool, 1)
	defer func() {
		if r := recover(); r != nil {
			log.Recover(r)
		}

		closeSig <- true
	}()

	handleMsgData := func(args []interface{}) (error) {
		if a.gate.Processor != nil {
			data := args[0].([]byte)
			msg, err := a.gate.Processor.Unmarshal(data)
			if err != nil {
				return err
			}

			err = a.gate.Processor.Route(msg, a)
			if err != nil {
				return err
			}
		}
		return nil
	}

/*	if a.chanRPC != nil {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Recover(r)
				}

				if a.gate.OnAgentDestroy != nil {
					a.gate.OnAgentDestroy(a)
				}
			}()

			//a.chanRPC.Register("handleMsgData", handleMsgData)

			if a.gate.OnAgentInit != nil {
				a.gate.OnAgentInit(a)
			}

			//a.skeleton.Run(closeSig)
		}()
	}*/

	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			//log.Debug("read message: %v", err)
			break
		}
		//等待切换module

		if a.chanRPC == nil {
			err = handleMsgData([]interface{}{data})
		} else {
			err = a.chanRPC.Call0("handleMsgData",a,data)
		}
		if err != nil {
			log.Debug("handle message: %v", err)
			break
		}
	}
}

func (a *agent) OnClose() {
	if a.chanRPC != nil {
		err := a.chanRPC.Call0("CloseAgent", a)
		if err != nil {
			log.Error("chanrpc error: %v", err)
		}
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

/*func (a *agent) Skeleton() *module.Skeleton {
	return a.skeleton
}*/

func (a *agent) ChanRPC() *chanrpc.Server {
	return a.chanRPC
}

func (a *agent) Gate()  *Gate {
	return a.gate
}
func (a* agent)SetChanRPC(rpc *chanrpc.Server){
	a.chanRPC=rpc
}