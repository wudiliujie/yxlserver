package center

import (
	"gameserver/center/internal"

	"gameserver/common"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
func RegisterRoomModule( module common.RoomModule){
	internal.RegisterRoomModule( module)
}