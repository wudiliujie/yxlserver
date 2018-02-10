package internal

import (
	"leaf/gate"
	"gameserver/common"
)
var (
	playerMap= map[int32]*Player{}
)
type Player struct {
	RoleId int32
	intAttr  map[int32]int64
	strAttr  map[int32]string
	Agent gate.Agent
	Module  common.RoomModule
}
func GetPlayer(roleId int32)(*Player){
	player, ok := playerMap[roleId]
	if ok  {
		return  player;
	}
	return  nil;
}
func OnPlayerLoginSuccess(roleId int32, module common.RoomModule){
	player:= GetPlayer(roleId)
	if player!=nil{
		player.Module=module;
	}else{
		player =new (Player)
		player.Module=module
		playerMap[roleId]= player
	}
}
func OnPlayerLogout(args[] interface{})error{

	roleId:=args[0].(int32)
	player:=GetPlayer(roleId)
	if player !=nil{
		//player.Agent=nil
	}
	delete(playerMap,roleId)
	return  nil
}

func  OnPlayerChangeModule(args[] interface{}){
	roleId := args[0].(int32)
	module:=args[1].(common.RoomModule)
	player:= GetPlayer(roleId)
	if(player!=nil){
		player.Module=module;
	}else{
		player =new (Player)
		player.Module=module
	}
}