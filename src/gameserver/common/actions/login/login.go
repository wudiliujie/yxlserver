package login

import (
	"common/proto"
	"gameserver/common/I"
	"gameserver/common/actions"
	"gameserver/common/playerManage"
	"gameserver/events"
	"leaf/log"
)

func init() {
	actions.RegisterPacket(proto.PCK_C2S_Login_ID, OnLogin)
}

func OnLogin(_player I.IPlayer, _pckMsg interface{}) {
	player := _player.(*playerManage.Player)
	pckMsg := _pckMsg.(*proto.C2S_Login)
	log.Debug("登录:%v>>>", player, pckMsg)
	//去数据库线程初始化数据
	player.AsynCall(events.DB_LoadRoleData, func(ret interface{}) {
		log.Debug("callback:%v", ret)
	})
}
