package internal

import (
	"leaf/module"
	"gameserver/base"
	"gameserver/gate"
	"gameserver/common"
	"math"
	"gameserver/center/players"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	PlayerManager = players.NewPlayerManager()
	roomModule = make(map[int]common.RoomModule)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	gate.Module.Gate.AgentChanRPC=ChanRPC
}

func (m *Module) OnDestroy() {

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