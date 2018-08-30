package playerManage

import (
	"gameserver/base"
	"gameserver/common/actions"
	"gameserver/events"
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

func Update(diff int64) {
	for _, v := range playerMap {
		v.Update(diff)
	}
}

func GetPlayer(roleId int32) *Player {
	return playerMap[roleId]
}
func NewPlayer(roleId int32) *Player {
	player := &Player{RoleId: roleId}
	player.queueMsg = queue.New()
	playerMap[roleId] = player
	return player
}

//添加网络消息
func (p *Player) AddNetMsg(msgId int32, pck interface{}) {
	//log.Debug("msgId:%v", msgId)
	info := &MsgInfo{msgId: msgId, pck: pck}
	p.queueMsg.Add(info)
	p.handleNetMsg()
}
func (p *Player) handleNetMsg() {

	if p.waitHandle <= 0 && p.queueMsg.Length() > 0 {
		msgInfo := p.queueMsg.Remove().(*MsgInfo)

		actions.ActionHandle(p, msgInfo.msgId, msgInfo.pck)
	}
}
func (p *Player) AsynCall(eventType events.EventType, args ...interface{}) {
	log.Debug("开启rpc")
	p.waitHandle = 1 //停止处理消息
	cb := args[len(args)-1].(func(ret interface{}))
	base.APP.AsynCall(base.DB.ChanRPCServer, eventType, func(ret interface{}, err error) {
		p.waitHandle = 0
		log.Debug("结束rpc:%v", err)
		cb(ret)
	})
}

func (p *Player) Update(diff int64) {
	p.handleNetMsg()

}
