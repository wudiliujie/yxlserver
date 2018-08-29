package playerManage

import (
	"gameserver/common/actions"
	"leaf/gate"
	"leaf/log"
	"leaf/util/queue"
)

type MsgInfo struct {
	msgId int32
	pck   interface{}
}

type Player struct {
	RoleId     int32
	agent      gate.Agent
	waitHandle int32        //等待处理消息时间
	queueMsg   *queue.Queue //消息队列

}

var playerMap = make(map[int32]*Player)

func GetPlayer(roleId int32) *Player {
	return playerMap[roleId]
}
func NewPlayer() *Player {
	player := &Player{}
	player.queueMsg = queue.New()
	return player
}

//添加网络消息
func (p *Player) AddNetMsg(msgId int32, pck interface{}) {
	log.Debug("msgId:%v", msgId)
	info := &MsgInfo{msgId: msgId, pck: pck}
	p.queueMsg.Add(info)
	p.handleNetMsg()
}
func (p *Player) handleNetMsg() {
	if p.waitHandle <= 0 && p.queueMsg.Length() > 0 {
		msgInfo := p.queueMsg.Peek().(*MsgInfo)
		actions.ActionHandle(p, msgInfo.msgId, msgInfo.pck)
	}
}
