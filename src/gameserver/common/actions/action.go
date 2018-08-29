package actions

import (
	"common/proto"
	"gameserver/common/I"
	"leaf/log"
)

var pckHandleMap = make(map[int32]func(player I.IPlayer, pckMsg interface{}))

func init() {
	pckHandleMap[int32(proto.PCK_C2S_Login_ID)] = OnLogin
}
func ActionHandle(player I.IPlayer, msgId int32, pck interface{}) {
	f, ok := pckHandleMap[msgId]
	if ok {
		f(player, pck)
	} else {
		log.Error("msg:%v 没有注册", msgId)
	}
}

func OnLogin(player I.IPlayer, pckMsg interface{}) {
	log.Debug("登录")
}
