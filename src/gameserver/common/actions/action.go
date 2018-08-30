package actions

import (
	"common/proto"
	"gameserver/common/I"
	"leaf/log"
)

var pckHandleMap = make(map[int32]func(player I.IPlayer, pckMsg interface{}))

func init() {
}
func RegisterPacket(msgId proto.PCK, f func(player I.IPlayer, pckMsg interface{})) {
	pckHandleMap[int32(msgId)] = f
}
func ActionHandle(player I.IPlayer, msgId int32, pck interface{}) {
	f, ok := pckHandleMap[msgId]
	if ok {
		f(player, pck)
	} else {
		log.Error("msg:%v 没有注册", msgId)
	}
}
