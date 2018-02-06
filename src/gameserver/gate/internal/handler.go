package internal

import (
	"common/proto"
	"leaf/gate"

	"gopkg.in/mgo.v2/bson"
	"fmt"
	"leaf/log"
)

func init() {
	//msg.Processor.SetHandler(&msg.C2S_Login{}, handleLogin)
	//msg.Processor.SetRouter(&msg.C2S_GetInfo{}, center.ChanRPC)
}

func onAgentInit(agent gate.Agent) {
	fmt.Print("客户端连接成功:%v",agent.RemoteAddr())
}

func onAgentDestroy(agent gate.Agent)() {
	fmt.Print("客户端断开连接:%v",agent.RemoteAddr())
/*	var accountId bson.ObjectId
	if val, ok := agent.UserData().(bson.ObjectId); ok {
		accountId = val
	}*/
/*	if accountId.Valid() {
		center.ChanRPC.Go("AccountOffline", accountId, agent)
	}*/
}
func handleLogin(args []interface{}){
	recvMsg := args[0].(*proto.C2S_Login)
	log.Debug("loing:%s>>>>%s",recvMsg.Name,recvMsg.Password)
	agent := args[1].(gate.Agent)
	accountid := bson.NewObjectId();
	agent.SetUserData(accountid)
	//center.ChanRPC.Go("CreatePlayer", accountid, agent)
}
func  handleGetInfo(args []interface{})  {
	log.Debug("获取信息");
}

