package internal

import (
	"common/errmsg"
	"consts"
	"gameserver/center"
	"gameserver/players"
)

type RoomData struct {
	Id       int32
	RoomType consts.ROOM_TYPE
	Players  map[int32]*players.Player
}

var (
	roomMap map[int32]*RoomData
)

func (m *Module) CreateRoom(roomId int32, roomType consts.ROOM_TYPE) *RoomData {
	room := &RoomData{Id: roomId, RoomType: roomType}
	room.Players = make(map[int32]*players.Player)
	m.roomMap[roomId] = room
	//通知主线程创建房间
	center.ChanRPC.Call0(consts.Center_Rpc_OnCreateRoom, m, roomId, roomType)
	return room
}
func (m *Module) GetRoom(roomId int32) *RoomData {
	room, ok := m.roomMap[roomId]
	if ok {
		return room
	}
	return nil
}

func (room *RoomData) AddPlayer(player * players.Player) int32 {
	room.Players[player.RoleId] = player
	player.RoomId= room.Id
	//room.OnRoomNumChange() //添加的时候，直接在主线程加1了
	if len(room.Players) >= 5 {
	}
	//log.Debug("人数大于5:%v", atomic.LoadInt32(&LoginSuccess))
	return errmsg.SYS_SUCCESS
}
func (room* RoomData) RemovePlayer(roleId int32){
	delete(room.Players,roleId)
	room.OnRoomNumChange(-1)
}


//当玩家进入房间
func (room *RoomData) OnRoomNumChange(num int32) {
	center.ChanRPC.Call0(consts.Center_Rpc_OnRoomNumChange, room.Id, num)
}
