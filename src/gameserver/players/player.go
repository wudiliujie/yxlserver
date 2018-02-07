package players

import (
	"leaf/gate"
	"gameserver/common"
	"common/proto"
	proto1 "github.com/golang/protobuf/proto"
	"leaf/log"
)

var (
	playerPool      = make(map[int]*Player)
)


type Player struct {
	AccountId int
	intAttr  map[int32]int64
	strAttr  map[int32]string
	Agent gate.Agent
	Module  common.RoomModule
}
func GetPlayer(accountid int)(*Player){
	player, ok := playerPool[accountid]
	if ok  {
		return  player;
	}
	return  nil;
}
func CreatePlayer(accountid int)(*Player){
	player, ok := playerPool[accountid]
	if ok  {
		return  nil;
	}

	player = new(Player);
	player.intAttr= make(map[int32]int64)
	player.strAttr= make(map[int32]string)
	player.AccountId= accountid;
	playerPool[accountid]= player;
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

