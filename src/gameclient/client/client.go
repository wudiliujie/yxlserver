package client

import (
	"leaf/log"
	"leaf/network"
	"net"
	"reflect"
	"sync/atomic"

	"common/proto"
	"crypto/md5"
	"fmt"
	"strconv"
	sync "sync"
)

var (
	Client      *Agent
	processor   network.Processor
	ClientSN    int32
	clientslock = sync.Mutex{}
	clients     = map[int32]*Agent{}
)

func GetClientSNNew() int32 {
	return atomic.AddInt32(&ClientSN, 1)
}
func newAgent(conn *network.TCPConn) network.Agent {
	c := new(Agent)
	c.conn = conn
	c.id = GetClientSNNew()
	//连接成功
	//发送登录
	name := strconv.Itoa(int(c.id))
	password := "111111"
	userData.AccountName = name
	hash := md5.Sum([]byte(password))
	strMd5 := fmt.Sprintf("%x", hash)
	msg := &proto.C2S_Login{Name: name, Password: strMd5}
	log.Debug("发送登录包")
	c.WriteMsg(msg)
	clientslock.Lock()
	defer clientslock.Unlock()
	clients[c.id] = c

	return c
}

type Agent struct {
	id       int32
	conn     *network.TCPConn
	userData interface{}
}

func (a *Agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Error("read message: %v", err)
			break
		}

		if processor != nil {
			msg, err := processor.Unmarshal(data)
			if err != nil {
				log.Error("unmarshal message error: %v", err)
				break
			}
			err = processor.Route(msg, a)
			if err != nil {
				log.Error("route message error: %v", err)
				break
			}
		}
	}
}

func (a *Agent) OnClose() {}

func (a *Agent) WriteMsg(msg interface{}) {
	if processor != nil {
		data, err := processor.Marshal(msg)
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

func (a *Agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *Agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *Agent) Close() {
	a.conn.Close()
}

func (a *Agent) Destroy() {
	a.conn.Destroy()
}

func (a *Agent) UserData() interface{} {
	return a.userData
}

func (a *Agent) SetUserData(data interface{}) {
	a.userData = data
}

func Init(p network.Processor) {
	processor = p
}

func Close() {
	if Client != nil {
		Client.Destroy()
		Client = nil
	}
}

func Start(addr string) *network.TCPClient {
	//Close()

	client := new(network.TCPClient)
	client.Addr = addr
	client.NewAgent = newAgent
	client.ConnNum = 1
	client.ConnectInterval = 3
	client.PendingWriteNum = 100
	client.Start()

	//log.Release("start connect to %v", addr)
	/*for {
		time.Sleep(time.Second)
		if Client != nil {
			break
		}
	}
	log.Release("connect %v success", addr)*/
	return client
}
