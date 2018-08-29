package players

import (
	"leaf/gate"
	"leaf/util"

	"common/errmsg"
	"common/proto"
	"consts"
	proto1 "github.com/golang/protobuf/proto"
	"leaf/log"
)

type Player struct {
	RoleId  int32
	intAttr map[int32]int64
	strAttr map[int32]string
	Agent   gate.Agent
	Module  common.RoomModule
	RoomId  int32
}

var (
	playerMap util.Map
)

func GetPlayer(roleId int32) *Player {
	_p := playerMap.Get(roleId)
	if _p != nil {
		return _p.(*Player)
	}
	return nil
}
func CreatePlayer(roleId int32, agent gate.Agent, m common.RoomModule) *Player {

	if playerMap.Get(roleId) != nil {
		log.Debug("创建角色的时候，用户已经存在%v", roleId)
		return nil
	}
	player := &Player{}
	player.intAttr = make(map[int32]int64)
	player.strAttr = make(map[int32]string)
	player.RoleId = roleId
	player.Agent = agent
	player.Module = m
	playerMap.Set(roleId, player)
	return player
}
func RemovePlayer(roleId int32) {
	playerMap.Del(roleId)
}
func GetPlayerNum() int32 {
	return int32(playerMap.Len())
}
func (player *Player) InitData(dbinfo *[]byte) (isNew bool) {
	isNew = false
	info := &proto.RoleDbInfo{}
	proto1.UnmarshalMerge(*dbinfo, info)
	for _, v := range info.IntAttr {
		player.intAttr[v.K] = v.V
	}
	for _, v := range info.StrAttr {
		player.strAttr[v.K] = v.V
	}
	if len(player.intAttr) == 0 {
		player.onRoleCreate()
		isNew = true
	}
	//playerManage.intAttr[1]=100;
	return
	//playerManage.status=1
}
func (player *Player) onRoleCreate() {
	player.intAttr[1] = int64(player.RoleId)
}
func (player *Player) GetSaveData() []byte {
	info := &proto.RoleDbInfo{}
	info.Roleid = 0
	for k, v := range player.intAttr {
		info.IntAttr = append(info.IntAttr, &proto.IntAttr{k, v})
	}
	for k, v := range player.strAttr {
		info.StrAttr = append(info.StrAttr, &proto.StrAttr{k, v})
	}
	data, err := proto1.Marshal(info)
	if err != nil {
		log.Error("序列化错误%v", err.Error())
	}
	return data
}
func (player *Player) SendMsg(msg interface{}) {
	player.Agent.WriteMsg(msg)
}
func (player *Player) ReplaceAgent(agent gate.Agent) {
	if player.Agent != nil {
		player.Agent.Close()
	}
	player.Agent = agent
}
func (player *Player) SendRoleInfo() {
	send := &proto.S2C_GetInfo{}
	send.IntAttr = append(send.IntAttr, &proto.IntAttr{int32(proto.PLAYER_ATTR_GOLD), 100})
	send.IntAttr = append(send.IntAttr, &proto.IntAttr{int32(proto.PLAYER_ATTR_DIAMOND), 1000})
	send.IntAttr = append(send.IntAttr, &proto.IntAttr{int32(proto.PLAYER_ATTR_LEVEL), 1})
	send.IntAttr = append(send.IntAttr, &proto.IntAttr{int32(proto.PLAYER_ATTR_EXP), 0})
	player.SendMsg(send)
}
func (player *Player) ChangeModule(old common.RoomModule, new common.RoomModule, roomId int32, cb func(e error)) int32 {
	new.GetChanRPC().Go(consts.Room_Rpc_EnterModule, player, roomId, cb)
	old.RemovePlayer(player.RoleId)
	return errmsg.SYS_ASYNC_WAIT
}

/*




func (m *Module) EnterRoom(playerManage *Player, roomType consts.ROOM_TYPE) (int32,int32){
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
		if playerManage.Module.GetId()==module.GetId(){ //本模块内
			//直接添加房间
			room:=m.GetRoom(roomId)
			if room==nil{
				log.Error("本模块房间不存在")
				return errmsg.SYS_ROOM_SYSTEM_ERROR,0
			}
			return room.AddPlayer(playerManage),roomId
		}else{ //切换模块
			ret:=playerManage.ChangeModule(m,module,roomId, func(e error){
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
		room.AddPlayer(playerManage)
	}
	return  errmsg.SYS_SUCCESS,roomId
}
*/
