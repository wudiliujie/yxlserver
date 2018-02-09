package client

import (
	"leaf/module"
	"gameclient/base"
	"time"
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
)

type ClientModule struct {
	*module.Skeleton
}

func (m *ClientModule) OnInit() {
	m.Skeleton = skeleton
	args:=[]interface{}{"a","b"}
	for i:=0;i<10000;i++{
		 go login(args)
		 time.Sleep(time.Microsecond*10)
	}
}

func (m *ClientModule) OnDestroy() {

}
