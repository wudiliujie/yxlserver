package client

import (

	"common/proto"
	"leaf/log"
	"common/msg"
	"common/errmsg"
)

func init() {
	msg.Processor.SetHandler(&proto.S2C_Login{}, handleLogin)
	msg.Processor.SetHandler(&proto.S2C_GetInfo{}, handleGetInfo)
}

func handleLogin(args []interface{}) {
	recvMsg := args[0].(*proto.S2C_Login)
	if recvMsg.Tag !=errmsg.SYS_SUCCESS {
		Close()
		log.Error("login is error: %v, please input login cmd", recvMsg.Tag)
		return
	}
	//
	log.Debug("登录成功，获取信息")
	getinfo:=proto.C2S_GetInfo{};
	log.Debug("登录成功，获取信息1")
	Client.WriteMsg(&getinfo);
	log.Debug("登录成功，获取信息2")

}
func handleGetInfo(args []interface{}){
	recvMsg := args[0].(*proto.S2C_GetInfo)
	log.Debug("获取信息返回")
	for _,v :=range recvMsg.IntAttr{
		log.Debug("K:%v V:%v",v.K,v.V)
	}
	log.Debug("发送进入房间")
	sendmsg:=&proto.C2S_EnterRoom{RoomType:0}
	Client.WriteMsg(sendmsg)

}

