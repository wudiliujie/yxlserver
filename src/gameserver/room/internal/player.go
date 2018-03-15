package internal

import (

	"gameserver/common"
	"common/proto"
	"leaf/log"
	"common/errmsg"
	"consts"
	"gameserver/center"
	"gameserver/players"
)



func (m *Module)GetPlayer(roleId int32)(*players.Player){
	player, ok := m.playerMap[roleId]
	if ok  {
		return  player;
	}
	return  nil;
}
//进入模块
func (m* Module)AddPlayer(player *players.Player){
	m.playerMap[player.RoleId]=player
	player.Agent.SetChanRPC(m.ChanRPC)
	m.AddClientCount(1)
}

func (m* Module)RemovePlayer(roleId int32){
	delete(m.playerMap, roleId)
	m.AddClientCount(-1)

}
//进入模块
func (m* Module)HandleEnterModule(args[] interface{})error{
	player:= args[0].(*players.Player)
	roomId:=args[1].(int32)

	m.AddPlayer(player)
	if roomId!=0{
		room:= m.GetRoom(roomId)
		if room!=nil{
			room.AddPlayer(player)
		}
		//通知玩家进入房间成功
		sendmsg:=&proto.S2C_EnterRoom{Tag:errmsg.SYS_SUCCESS,RoomId:roomId}
		player.SendMsg(sendmsg)
	}
	return  nil
}









func (m *Module) EnterRoom(player *players.Player, roomType consts.ROOM_TYPE) (int32,int32){
	args,err:=center.ChanRPC.Call1(consts.Center_Rpc_GetBestRoomId,roomType)
	if err!=nil {
		log.Error("获取房间系统错误:%v",err)
		return errmsg.SYS_ROOM_SYSTEM_ERROR,0
	}
	roomId:= args.(int32)
	if roomId!=0{
		//进入
		//获取房间所在的module
		args,err := center.ChanRPC.Call1(consts.Center_Rpc_GetModuleByRoomId,roomId)
		if err!=nil{
			log.Error("获取房间所在的MODULE系统错误:%v>>.%v",roomId,err)
			return errmsg.SYS_ROOM_SYSTEM_ERROR,0
		}
		module:=args.(common.RoomModule)
		if player.Module.GetId()==module.GetId(){ //本模块内
			//直接添加房间
			room:=m.GetRoom(roomId)
			if room==nil{
				log.Error("本模块房间不存在")
				return errmsg.SYS_ROOM_SYSTEM_ERROR,0
			}
			return room.AddPlayer(player),roomId
		}else{ //切换模块
			ret:=player.ChangeModule(m,module,roomId, func(e error){
				if e!=nil{
					log.Error("切换房间进入模块失败:%v",e)
				}
			})
			return  ret,0

		}
	}else {
		//创建新的房间
		//log.Debug("创建新房间");
		//获取新房间编号
		args,err:=center.ChanRPC.Call1(consts.Center_Rpc_GetNewRoomId)
		if err!=nil {
			log.Error("获取房间系统错误:%v",err)
			return errmsg.SYS_ROOM_SYSTEM_ERROR,0
		}
		roomId=args.(int32)
		room:=m.CreateRoom(roomId,roomType)
		if room==nil{
			log.Error("创建房间为空")
			return errmsg.SYS_ROOM_SYSTEM_ERROR,0
		}
		room.AddPlayer(player)
	}
	return  errmsg.SYS_SUCCESS,roomId
}