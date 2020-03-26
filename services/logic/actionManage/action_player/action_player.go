package action_player

import (
	"github.com/wudiliujie/common/network"
	"yxlserver/services/I"
	"yxlserver/services/logic/actionManage"
	"yxlserver/services/msg"
)

func init() {
	actionManage.RegisterPlayerPacket(msg.PCK_C2SLoginEnd_ID, OnC2SLoginEnd)
}
func OnC2SLoginEnd(player I.IPlayer, _ network.IMessage) {
	player.OnAreaClose()
}
