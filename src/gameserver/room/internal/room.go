package internal

import (
	"common/errmsg"
	"consts"
	"gameserver/center"
	"gameserver/players"
	"gameserver/units"
	"leaf/util"
)

const(
	RoomState_Wait    = 0 //等待中
	RoomState_Gameing = 1//游戏中
	RoomState_GameOver = 2//游戏结束
)
//房子
type House struct {
	unit *units.Unit
	gold    int32 //价值
	issell  int32 //是否出售
}
func CreateHouse()*House{
	h:=&House{}
	h.gold=100;
	h.unit=nil
	return  h
}
//获取指定等级的房间购买价格
func (h* House)GetBuyGold(circle int32)int32{
	return  int32(circle*100)
}
func (h* House)Buy(circle int32){
	 h.gold= h.GetBuyGold(circle)
}
func (h* House)GetPrice()int32{
	return  h.gold
}
//获取升级的金币
func (h *House) GeteLevelGold(circle int32)int32{
	return  h.GetBuyGold(circle)-h.gold
}
func (h* House)LevelUp(){
	h.gold= h.gold+100
}
//获取税收
func (h *House) GetTaxes()int32{
	return   h.gold
}


type RoomData struct {
	Id       int32
	RoomType consts.ROOM_TYPE
	Players  map[int32]*players.Player
	Units  map[int32]*units.Unit
	RoomState	int32	//房间状态，0准备中，1游戏中，2游戏结束
	Houses   [consts.MAX_HOUSE_NUM+1]*House
	debugTimer util.CTimer
}

var (
	roomMap map[int32]*RoomData
)

func (m *Module) CreateRoom(roomId int32, roomType consts.ROOM_TYPE) *RoomData {
	room := &RoomData{Id: roomId, RoomType: roomType}
	room.Players = make(map[int32]*players.Player)
	room.Units=make(map[int32]*units.Unit)
	room.debugTimer.BeginTimer(1000,0)
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
