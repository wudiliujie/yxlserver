package internal

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"leaf/chanrpc"
	"leaf/log"
	"gameserver/players"
	"gameserver/common"
	"common/msg"
	"common/proto"
)

func RegisterHandler(chanRPC *chanrpc.Server,module *Module) {
	chanRPC.Register("EnterModule", EnterModule)
	chanRPC.Register("LeaveRoom", LeaveRoom)
	chanRPC.Register("SendMsg", SendMsg)

	chanRPC.Register("CloseAgent", module.CloseAgent)
	chanRPC.Register("handleMsgData",common.HandleMsgData)
	msg.Processor.SetHandler(&proto.C2S_GetInfo{}, handleGetInfo)
}
func handleGetInfo(args[] interface{}){
	log.Debug("获取用户信息")
}
func EnterModule(args []interface{}) (error) {
	log.Debug("进入模块")
	module := args[0].(*Module)
	player := args[1].(* players.Player)
	module.PlayerEnter(player);
	return  nil;
}

func LeaveRoom(args []interface{}) error {
	module := args[0].(*Module)
	roomName := args[1].(string)
	userId := args[2].(bson.ObjectId)

	roomInfo := module.GetRoomInfo(roomName)
	if roomInfo == nil {
		return errors.New("this room is not exist")
	}

	if !roomInfo.LeaveRoom(userId) {
		return errors.New("you is not in this room")
	}
	return nil
}

func SendMsg(args []interface{}) error {
	module := args[0].(*Module)
	roomName := args[1].(string)
	//userId := args[2].(bson.ObjectId)
	//msgContent := args[3].([]byte)

	roomInfo := module.GetRoomInfo(roomName)
	if roomInfo == nil {
		return errors.New("this room is not exist")
	}
	return  nil;
}

