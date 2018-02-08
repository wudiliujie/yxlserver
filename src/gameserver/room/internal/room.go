package internal

import (
	"consts"
	"common/errmsg"
	"gameserver/center"
)

type RoomData struct {
	Id int32
	RoomType consts.ROOM_TYPE
	Players map[int32]*Player
}

var (
	roomMap map[int32]* RoomData
)

func  (m *Module)CreateRoom(roomId int32 ,roomType consts.ROOM_TYPE)*RoomData{
	room:=&RoomData{Id:roomId,RoomType:roomType};
	m.roomMap[roomId]=room
	//通知主线程创建房间
	center.ChanRPC.Call0(consts.Center_Rpc_OnCreateRoom,m,roomId,roomType)
	return  room
}
func (room* RoomData) AddPlayer(player *Player)int32{
	room.Players[player.RoleId]= player;
	room.onPlayerEnter(player)
	return  errmsg.SYS_SUCCESS
}
//当玩家进入房间
func (room* RoomData)onPlayerEnter(player* Player){
	center.ChanRPC.Call0(consts.Center_Rpc_OnRoomNumChange,room.Id,len(room.Players))
}