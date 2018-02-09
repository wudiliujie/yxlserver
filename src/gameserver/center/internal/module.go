package internal

import (
	"leaf/module"
	"gameserver/base"
	"gameserver/gate"
	"gameserver/common"
	"math"
	"time"
	"leaf/log"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	roomModule = make(map[int]common.RoomModule)
	TestCount=int32(0)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	gate.Module.Gate.AgentChanRPC=ChanRPC
	m.Skeleton.AfterFunc(time.Second*30, m.showInfo)
}

func (m *Module) OnDestroy() {

}
func (m* Module) showInfo(){
	log.Debug("%v:当前人数：%v, room：%v","center",len(playerMap),len(roomMap))
	m.Skeleton.AfterFunc(time.Second*30, m.showInfo)
}
func RegisterRoomModule( module common.RoomModule){
	roomModule[module.GetId()] = module
}
func  GetBestModule() (module common.RoomModule ,ok bool) {
	var minCount int32 = math.MaxInt32
	ok=false
	for _, _module := range roomModule {
		count := _module.GetClientCount()
		if count < minCount {
			module = _module
			minCount = count
			ok=true
		}
	}
	return
}