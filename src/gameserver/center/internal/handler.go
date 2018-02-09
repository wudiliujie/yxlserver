package internal

import (
	"common/msg"
	"common/proto"
	"gameserver/common"
	"gameserver/conf"
	"leaf/gate"
	"common/errmsg"
	"consts"
	"gameserver/db"
	"leaf/log"
)

var (
	routeMoreMap = map[interface{}]interface{}{}
)

func handleRpc(id interface{}, f interface{}, fType int) {
	//cluster.SetRoute(id, ChanRPC)
	ChanRPC.Register(id, f)
}

func init() {
	ChanRPC.Register("handleMsgData", common.HandleMsgData)
	ChanRPC.Register("NewAgent", rpcNewAgent)
	ChanRPC.Register("CloseAgent", rpcCloseAgent)
	ChanRPC.Register(consts.Center_Rpc_GetBestRoomId, getBestRoomId)
	ChanRPC.Register(consts.Center_Rpc_GetModuleByRoomId, getModuleByRoomId)
	ChanRPC.Register(consts.Center_Rpc_GetNewRoomId, getNewRoomId)
	ChanRPC.Register(consts.Center_Rpc_OnCreateRoom, onCreateRoom)
	ChanRPC.Register(consts.Center_Rpc_OnRoomNumChange, onRoomNumChange)
	ChanRPC.Register(consts.Center_Rpc_OnPlayerLogout, OnPlayerLogout)


	msg.Processor.SetHandler(&proto.C2S_Login{}, handleLogin)
}
func rpcNewAgent(args []interface{}) {
	//a := args[0].(gate.Agent)
	//log.Debug("连接成功%v-->%v", a.RemoteAddr(),atomic.LoadInt32(&TestCount))
}
func rpcCloseAgent(args []interface{}) error {
	//a := args[0].(gate.Agent)
	//log.Debug("断开连接%v", a.RemoteAddr())
	return  nil
}

func handleLogin(args []interface{}) {


	recvMsg := args[0].(*proto.C2S_Login)
	//log.Debug("loing:%s>>>>%s", recvMsg.Name, recvMsg.Password)
	agent := args[1].(gate.Agent)
	accountid := int32(0)
	skeleton.Go(func() {
		accountid = db.GetAccountId(recvMsg.Name, recvMsg.Password)
	}, func() {

		agent.SetUserData(accountid)
		sendMsg := &proto.S2C_Login{Tag: errmsg.SYS_LOGIN_NOACCOUNT}
		sendMsg.Tag = errmsg.SYS_SUCCESS

		player := GetPlayer(accountid)
		if player != nil { //
			//已经在线,替换agent
			//player.ReplaceAgent(agent);
			log.Debug("重复登录:%v",accountid)
			sendMsg.Tag = errmsg.SYS_LOGIN_ALREAD_ONLINE
			agent.WriteMsg(sendMsg)
			return

		}
		roomModule, ok := GetBestModule()
		if !ok {
			sendMsg.Tag = errmsg.SYS_LOGIN_NO_MODULE
			agent.WriteMsg(sendMsg)
			return
		}
		skeleton.AsynCall(roomModule.GetChanRPC(), consts.Room_Rpc_LoginModule, accountid, agent, func(e error) {
			//log.Debug("登录结果:%v", e)
			//通知主线程登录成功
			OnPlayerLoginSuccess(accountid, roomModule)
		})
	})

}
func GetChatInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{common.GetClientCount(), conf.Server.ListenAddr}, nil
}
