package player

import (
	"github.com/wudiliujie/common/gate"
	"yxlserver/services/I"
	"yxlserver/services/consts"
	"yxlserver/services/logic/playerManage"
	"yxlserver/services/logic/playerManage/unit"
	"yxlserver/services/msg"
)

//玩家
type Player struct {
	unit.Unit
	state consts.PlayerState
	agent gate.Agent
}

func init() {
	playerManage.CreatePlayerFunc = func(userId int64) I.IPlayer {
		p := &Player{}
		p.Id = userId
		p.Init()
		return p
	}
}
func (p *Player) GetState() consts.PlayerState {
	return p.state
}
func (p *Player) SetState(state consts.PlayerState) {
	p.state = state
}
func (p *Player) InitUserData(data *msg.RoleDbData) {

}
func (p *Player) GetAgent() gate.Agent {
	return p.agent
}
func (p *Player) ChangeAgent(agent gate.Agent) {
	p.agent = agent
	agent.SetUserData(p.Id)
	p.SetState(consts.PlayerState_OnLine)
}

func (p *Player) OnAreaClose() {

}
func (p *Player) OnLoginSuccess() {

}
