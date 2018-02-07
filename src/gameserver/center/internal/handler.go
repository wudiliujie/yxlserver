package internal

import (
	"gameserver/conf"
	"gameserver/common"
	"leaf/gate"
	"leaf/log"
	"common/msg"
	"common/proto"

	"common/errmsg"
	"math/rand"
	"consts"
	"gameserver/center/rooms"
)

var (
	routeMoreMap = map[interface{}]interface{}{}
)

func handleRpc(id interface{}, f interface{}, fType int) {
	//cluster.SetRoute(id, ChanRPC)
	ChanRPC.Register(id, f)
}

func init() {
	ChanRPC.Register("handleMsgData",common.HandleMsgData)
	ChanRPC.Register("NewAgent",rpcNewAgent)
	ChanRPC.Register("CloseAgent",rpcCloseAgent)
	ChanRPC.Register(consts.Center_Rpc_GetBestRoomId,rooms.GetBestRoomId)
	ChanRPC.Register(consts.Center_Rpc_GetModuleByRoomId,rooms.GetModuleByRoomId)
	msg.Processor.SetHandler(&proto.C2S_Login{}, handleLogin)
}
func rpcNewAgent(args []interface {}) {
	a:= args[0].(gate.Agent)
	log.Debug("连接成功%v",a.RemoteAddr())
}
func rpcCloseAgent(args []interface {}) {
	a:= args[0].(gate.Agent)
	log.Debug("断开连接%v",a.RemoteAddr())
}


func handleLogin(args []interface{}){
	log.Debug("登录消息");
	recvMsg := args[0].(*proto.C2S_Login)
	log.Debug("loing:%s>>>>%s",recvMsg.Name,recvMsg.Password)
	agent := args[1].(gate.Agent)
	accountid := rand.Int31n(10000000)
	agent.SetUserData(accountid)
	sendMsg := &proto.S2C_Login{Tag:errmsg.SYS_LOGIN_NOACCOUNT}
	sendMsg.Tag=errmsg.SYS_SUCCESS

	player:= PlayerManager.GetPlayer(accountid);
	if player!=nil{ //
		//已经在线,替换agent
		//player.ReplaceAgent(agent);
		sendMsg.Tag=errmsg.SYS_LOGIN_ALREAD_ONLINE
		agent.WriteMsg(sendMsg)
		return

	}
	 roomModule ,ok:= GetBestModule()
	 if !ok{
		 sendMsg.Tag=errmsg.SYS_LOGIN_ALREAD_ONLINE
		 agent.WriteMsg(sendMsg)
		 return
	 }
	skeleton.AsynCall(roomModule.GetChanRPC(),consts.Room_Rpc_LoginModule,accountid,agent, func(error) {
	 	log.Debug("登录成功")
		//通知主线程登录成功
		PlayerManager.OnPlayerLoginSuccess(accountid,roomModule)
	 })

}
func GetChatInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{common.GetClientCount(), conf.Server.ListenAddr}, nil
}

