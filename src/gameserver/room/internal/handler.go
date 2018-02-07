package internal

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"leaf/chanrpc"
	"leaf/log"
	"gameserver/room/players"
	"common/msg"
	"common/proto"
	"consts"
)

func RegisterHandler(chanRPC *chanrpc.Server,module *Module) {

	chanRPC.Register("CloseAgent", module.CloseAgent)
	chanRPC.Register("handleMsgData",module.HandleMsgData)
	chanRPC.Register(consts.Room_Rpc_LoginModule,module.HandleLoginModule)
	msg.Processor.SetHandler(&proto.C2S_GetInfo{}, handleGetInfo)
	msg.Processor.SetHandler(&proto.C2S_EnterRoom{}, handleEnterRoom)


}
func handleGetInfo(args[] interface{}){
	log.Debug("获取用户信息")
	//recvMsg:= args[0].(*proto.C2S_GetInfo)
	player:=args[1].(*players.Player)
	player.SendRoleInfo()

}
//进入房间
func handleEnterRoom(args[] interface{}){
	recvMsg:= args[0].(*proto.C2S_EnterRoom)
	player:=args[1].(*players.Player)
	sendmsg:=&proto.S2C_EnterRoom{}
	sendmsg.Tag= player.EnterRoom(consts.ROOM_TYPE(recvMsg.RoomType))
	player.SendMsg(sendmsg)
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

