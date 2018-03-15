package client

import (
	"leaf/module"
	"gameclient/base"
	"time"
	"common/proto"
	"leaf/log"
	"gameclient/conf"
)

var (
	skeleton = base.NewSkeleton()
	Module   = new(ClientModule)
	ChanRPC  = skeleton.ChanRPCServer

	LoginCount=int32(0)
	LoginCount_Fail=int32(0)
	GetInfoCount=int32(0)
	GetInfoCount_Fail=int32(0)
	EnterRoomCount=int32(0)
	EnterRoomCount_Fail=int32(0)
	NUM=10000
)

type ClientModule struct {
	*module.Skeleton
}

func (m *ClientModule) OnInit() {
	m.Skeleton = skeleton
	args:=[]interface{}{"a","b"}
	for i:=0;i<conf.Client.ClientNum;i++{
		 go login(args)
		 time.Sleep(time.Microsecond*10)
	}
	skeleton.AfterFunc(time.Millisecond*500,SendInfo)
}
func SendInfo(){
	clientslock.Lock()
	defer clientslock.Unlock()
	if len(clients)>=conf.Client.ClientNum{
		log.Debug("发送信息")
		for _,v:=range clients{
			sendmsg:=&proto.C2S_Fire{Angle:100}
			v.WriteMsg(sendmsg)
		}
	}

	skeleton.AfterFunc(time.Millisecond*200,SendInfo)
}

func (m *ClientModule) OnDestroy() {

}
