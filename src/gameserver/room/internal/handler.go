package internal

import (
	"leaf/chanrpc"
	"common/msg"
	"common/proto"
	"consts"
	"sync/atomic"
	"leaf/log"
)

func RegisterHandler(chanRPC *chanrpc.Server,module *Module) {

	chanRPC.Register("CloseAgent", module.CloseAgent)
	chanRPC.Register("handleMsgData",module.HandleMsgData)
	chanRPC.Register(consts.Room_Rpc_LoginModule,module.HandleLoginModule)
	chanRPC.Register(consts.Room_Rpc_EnterModule,module.HandleEnterModule)
	chanRPC.Register(consts.Room_Rpc_EnterModuleRoom,module.EnterModuleRoom)

	msg.Processor.SetHandler(&proto.C2S_GetInfo{}, handleGetInfo)
	msg.Processor.SetHandler(&proto.C2S_EnterRoom{}, handleEnterRoom)


}
func handleGetInfo(args[] interface{}){
	//log.Debug("获取用户信息")
	//recvMsg:= args[0].(*proto.C2S_GetInfo)
	hArgs:=args[1].(*HandleArgs)
	hArgs.P.SendRoleInfo()

}
//进入房间
func handleEnterRoom(args[] interface{}){
	//log.Debug("登录消息")

	log.Debug("登录房间%v",	atomic.AddInt32(&RoomTestCount,1))
	//log.Debug("收到进入房间请求")
	recvMsg:= args[0].(*proto.C2S_EnterRoom)
	hArgs:=args[1].(*HandleArgs)
	sendMsg:=&proto.S2C_EnterRoom{}
	sendMsg.Tag,sendMsg.RoomId= hArgs.M.EnterRoom(hArgs.P, consts.ROOM_TYPE(recvMsg.RoomType))

	hArgs.P.SendMsg(sendMsg)
}