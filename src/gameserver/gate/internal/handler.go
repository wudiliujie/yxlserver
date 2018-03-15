package internal

import (
	"common/errmsg"
	"common/msg"
	"common/proto"
	"consts"
	"fmt"
	"gameserver/center"
	"gameserver/db"
	"gameserver/players"
	"leaf/gate"
	"leaf/log"
	"gameserver/common"
)

func init() {
	//msg.Processor.SetHandler(&msg.C2S_Login{}, handleLogin)
	//msg.Processor.SetRouter(&msg.C2S_GetInfo{}, center.ChanRPC)
	msg.Processor.SetHandler(&proto.C2S_Login{}, handleLogin)
}

func onAgentInit(agent gate.Agent) {
	fmt.Print("客户端连接成功:%v", agent.RemoteAddr())
}

func onAgentDestroy(agent gate.Agent) {
	fmt.Print("客户端断开连接:%v", agent.RemoteAddr())
	/*	var accountId bson.ObjectId
		if val, ok := agent.UserData().(bson.ObjectId); ok {
			accountId = val
		}*/
	/*	if accountId.Valid() {
		center.ChanRPC.Go("AccountOffline", accountId, agent)
	}*/
}
func handleGetInfo(args []interface{}) {
	log.Debug("获取信息")
}
func handleLogin(args []interface{}) {

	recvMsg := args[0].(*proto.C2S_Login)
	//log.Debug("loing:%s>>>>%s", recvMsg.Name, recvMsg.Password)
	agent := args[1].(gate.Agent)
	accountid := int32(0)

	accountid = db.GetAccountId(recvMsg.Name, recvMsg.Password)
	//agent.SetUserData(accountid)
	sendMsg := &proto.S2C_Login{Tag: errmsg.SYS_LOGIN_NOACCOUNT}
	sendMsg.Tag = errmsg.SYS_SUCCESS

	/*	player := GetPlayer(accountid)
		if player != nil { //
			//已经在线,替换agent
			//player.ReplaceAgent(agent);
			log.Debug("重复登录:%v",accountid)
			sendMsg.Tag = errmsg.SYS_LOGIN_ALREAD_ONLINE
			agent.WriteMsg(sendMsg)
			return

		}*/

	var data []byte
	db.ReadRoleInfo(accountid, data)

	_roomModule, err := center.ChanRPC.Call1(consts.Center_Rpc_GetBestModule)
	if err != nil {
		sendMsg.Tag = errmsg.SYS_LOGIN_NO_MODULE
		agent.WriteMsg(sendMsg)
		return
	}

	roomModule := _roomModule.(common.RoomModule)
	player := players.CreatePlayer(accountid,agent,roomModule)
	player.InitData(&data)

	err = roomModule.GetChanRPC().Call0(consts.Room_Rpc_LoginModule, accountid, agent)
	if err != nil {
		sendMsg.Tag = errmsg.SYS_LOGIN_NO_MODULE
		agent.WriteMsg(sendMsg)
		return
		//发送登录成功
	}
	center.ChanRPC.Call0(consts.Center_Rpc_PlayerLoginSuccess, accountid, roomModule)
	agent.SetUserData(accountid)
	sendmsg := &proto.S2C_Login{Tag: errmsg.SYS_SUCCESS}
	agent.WriteMsg(sendmsg)
	//log.Debug("登录成功")
}
