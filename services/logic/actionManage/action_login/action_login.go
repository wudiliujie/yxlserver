package action_login

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/network"
	"yxlserver/services/I"
	"yxlserver/services/consts"
	"yxlserver/services/errmsg"
	"yxlserver/services/logic/areaManage"
	"yxlserver/services/logic/playerManage"
	"yxlserver/services/util"

	"yxlserver/services/logic/actionManage"
	"yxlserver/services/msg"
)

//
func init() {
	actionManage.RegisterAgentPacket(msg.PCK_C2SLogin_ID, OnC2SLogin)
}

func OnC2SLogin(agent gate.Agent, _pckMsg network.IMessage) {
	//验证签名是否正确
	recvMsg := _pckMsg.(*msg.C2SLogin)
	//todo 这里需要验证签名是否超时
	sign := util.GetLoginSign(recvMsg.UserId, recvMsg.LoginTime)
	if sign != recvMsg.Sign {
		sendMsg := &msg.S2CLogin{}
		sendMsg.Tag = errmsg.SYS_签名错误
		agent.WriteMsg(0, sendMsg)
		return
	}
	playerManage.SignPlayerAsyncFunc(recvMsg.UserId, loadUserCallBack, agent, recvMsg.Sign)
}
func loadUserCallBack(player I.IPlayer, args ...interface{}) {
	agent := args[0].(gate.Agent)
	sign := args[1].(string)
	sendMsg := &msg.S2CLogin{}
	if player.GetState() == consts.PlayerState_Error {
		log.Error("LoadData:%v", player.GetId())
		sendMsg.Tag = errmsg.SYS_读取数据错误
		agent.WriteMsg(0, sendMsg)
		agent.Close()
	} else {
		//判断当前区服是否开启
		if areaManage.GetAreaState(player.GetInt32Value(consts.AttrType_LoginAreaId)) == consts.ServerState_Close {
			sendMsg.Tag = errmsg.SYS_服务器维护中
			agent.WriteMsg(0, sendMsg)
			agent.Close()
		}
		ChangeAgent(player, agent)
		//设置token，做断线重连用
		player.SetStrValue(consts.AttrType_LoginSign, sign)
	}
	agent.WriteMsg(0, sendMsg)
	agent.Close()
}

//发生顶号时，将老用户数据拷贝并剔除
func ChangeAgent(player I.IPlayer, newAgent gate.Agent) {
	//老用户踢出
	if util.IsNil(player.GetAgent()) == false {
		log.Debug("OLD:%v", player.GetAgent().RemoteAddr())
		killMsg := &msg.S2CKill{}
		killMsg.Tag = errmsg.SYS_帐号在其他地方登录
		player.GetAgent().WriteMsg(0, killMsg)
		//用户互顶，修改用户的离线时间
		//player.SendPck(killMsg)
		player.GetAgent().SetUserData(int64(0))
		player.GetAgent().Close()
	}
	//用户在30秒内登录成功，不要执行清理缓存操作 2019-07-30 刘杰
	//用户数据更换新的agent
	player.ChangeAgent(newAgent)
}
