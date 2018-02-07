package room

import (
	"gameserver/room/internal"
	"gameserver/conf"
	"leaf/module"
	"math"
	"leaf/chanrpc"
	"gameserver/center"
)

var (
	modules = []*internal.Module{}
)

type Module interface {
	GetChanRPC() *chanrpc.Server
}

func CreateModules() []module.Module {
	results := []module.Module{}
	for i := 0; i < conf.Server.RoomModuleCount; i++ {
		module := internal.NewModule(i)
		modules = append(modules, module)
		results = append(results, module)
		center.RegisterRoomModule(module)
	}
	return results
}

func GetBestModule() (module *internal.Module) {
	var minCount int32 = math.MaxInt32
	for _, _module := range modules {
		count := _module.GetClientCount()
		if count < minCount {
			module = _module
			minCount = count
		}
	}
	return
}