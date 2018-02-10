package internal

import (
	"leaf/module"
	"gameserver/base"
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

	m.Skeleton.AfterFunc(time.Second*30, m.showInfo)
}

func (m *Module) OnDestroy() {
 log.Debug("销毁centerModule")
}
func (m* Module) showInfo(){
	log.Debug("%v:当前人数：%v, room：%v","center",len(playerMap),len(roomMap))
	m.Skeleton.AfterFunc(time.Second*30, m.showInfo)

	destroyRooms := []int32{}
	for i,v:=range  roomMap{
		if v.Num<=0{
			destroyRooms=append(destroyRooms,i)
		}
	}
	if len(destroyRooms) > 0 {
		for _, idx := range destroyRooms {
			delete(roomMap, idx)
		}
		//cluster.Go("world", "DestroyRoom", destroyRooms)
		log.Debug("%v rooms is destroy in %v", destroyRooms, "center")
	}
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