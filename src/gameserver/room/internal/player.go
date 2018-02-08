package internal

import (
	"leaf/gate"
	"gameserver/common"
	"common/proto"
	proto1 "github.com/golang/protobuf/proto"
	"leaf/log"
	"common/errmsg"
	"consts"
	"gameserver/center"
)

var (

)


type Player struct {
	RoleId int32
	intAttr  map[int32]int64
	strAttr  map[int32]string
	Agent gate.Agent
	Module  common.RoomModule
}
func (m *Module)GetPlayer(roleId int32)(*Player){
	player, ok := m.playerMap[roleId]
	if ok  {
		return  player;
	}
	return  nil;
}
func (m *Module)CreatePlayer(roleId int32,agent gate.Agent)(*Player){
	player, ok :=  m.playerMap[roleId]
	if ok  {
		return  nil;
	}
	player = new(Player);
	player.intAttr= make(map[int32]int64)
	player.strAttr= make(map[int32]string)
	player.RoleId= roleId;
	player.Agent=agent
	m.playerMap[roleId]= player;
	return  player;
}

func (player* Player)ReplaceAgent(agent gate.Agent){
	if player.Agent!=nil{
		player.Agent.Close()
	}
	player.Agent= agent
}
func(player* Player) InitData(dbinfo *[]byte){
	info :=&proto.RoleDbInfo{}
	proto1.UnmarshalMerge(*dbinfo,info);
	for _,v :=range  info.IntAttr{
		player.intAttr[v.K]=v.V;
	}
	for _,v :=range info.StrAttr{
		player.strAttr[v.K]=v.V;
	}
	if len(player.intAttr)==0{
		player.onRoleCreate();
	}
	player.intAttr[1]=100;
	//player.status=1
}
func (player* Player) GetSaveData()*[]byte{
	info :=&proto.RoleDbInfo{}
	info.Roleid= 0;
	for k,v:=range player.intAttr{
		info.IntAttr= append(info.IntAttr,&proto.IntAttr{k,v})
	}
	for k,v:=range player.strAttr{
		info.StrAttr= append(info.StrAttr,&proto.StrAttr{k,v})
	}
	 data,err:= proto1.Marshal(info);
	 if(err!= nil){
	 	log.Error("序列化错误%v",err.Error())
	 }
	 return  &data;
}
func (player* Player) onRoleCreate(){
	player.intAttr[1]=100;
}
func (player* Player)SendRoleInfo(){
	send:= &proto.S2C_GetInfo{ }
	send.IntAttr= append(send.IntAttr,&proto.IntAttr{int32(proto.PLAYER_ATTR_GOLD),100})
	send.IntAttr= append(send.IntAttr,&proto.IntAttr{int32(proto.PLAYER_ATTR_DIAMOND),1000})
	send.IntAttr= append(send.IntAttr,&proto.IntAttr{int32(proto.PLAYER_ATTR_LEVEL),1})
	send.IntAttr= append(send.IntAttr,&proto.IntAttr{int32(proto.PLAYER_ATTR_EXP),0})
	player.SendMsg(send);
}
func (player* Player) SendMsg(msg interface{} ){
	player.Agent.WriteMsg(msg)
}
func (m *Module) EnterRoom(player *Player, roomType consts.ROOM_TYPE) (int32,int32){
	args,err:=center.ChanRPC.Call1(consts.Center_Rpc_GetBestRoomId,roomType)
	if err!=nil {
		log.Error("获取房间系统错误:%v",err)
		return errmsg.SYS_ROOM_SYSTEM_ERROR,0
	}
	roomId:= args.(int32)
	if roomId!=0{
		//进入
		//获取房间所在的module
		module,err := center.ChanRPC.Call1(consts.Center_Rpc_GetModuleByRoomId,roomId)
		if err!=nil{
			log.Error("获取房间所在的MODULE系统错误:%v>>.%v",roomId,err)
			return errmsg.SYS_ROOM_SYSTEM_ERROR,0
		}
		if player.Module==module{ //本模块内


		}else{ //切换模块

		}

	}else {
		//创建新的房间
		log.Debug("创建新房间");
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
	}
	return  errmsg.SYS_SUCCESS,roomId
}
