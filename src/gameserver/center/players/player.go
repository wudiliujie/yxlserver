package players

import (
	"leaf/gate"
	"gameserver/common"
)


type PlayerManager struct {
	playerPool map[int32]*Player
}

func  NewPlayerManager()*PlayerManager  {
	var result = new(PlayerManager)
	result.playerPool= make(map[int32]*Player)
	return  result
}
func (p *PlayerManager)Init(){
	p.playerPool= make(map[int32]*Player)
}
type Player struct {
	RoleId int32
	intAttr  map[int32]int64
	strAttr  map[int32]string
	Agent gate.Agent
	Module  common.RoomModule
}
func (p *PlayerManager)GetPlayer(roleId int32)(*Player){
	player, ok := p.playerPool[roleId]
	if ok  {
		return  player;
	}
	return  nil;
}
func (p* PlayerManager) OnPlayerLoginSuccess(roleId int32, module common.RoomModule){

	player:= p.GetPlayer(roleId)
	if player!=nil{
		player.Module=module;
	}else{
		player =new (Player)
		player.Module=module
		p.playerPool[roleId]= player
	}
}
func (p* PlayerManager) OnPlayerChangeModule(args[] interface{}){
	roleId := args[0].(int32)
	module:=args[1].(common.RoomModule)
	player:= p.GetPlayer(roleId)
	if(player!=nil){
		player.Module=module;
	}else{
		player =new (Player)
		player.Module=module
	}
}