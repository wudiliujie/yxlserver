package I

import (
	"github.com/wudiliujie/common/gate"
	"yxlserver/services/consts"
	"yxlserver/services/msg"
)

type IPlayer interface {
	IUnit
	OnAreaClose()
	OnLoginSuccess()
	GetState() consts.PlayerState
	SetState(state consts.PlayerState)
	InitUserData(data *msg.RoleDbData)
	GetAgent() gate.Agent
	ChangeAgent(agent gate.Agent)
}
