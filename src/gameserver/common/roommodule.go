package common

import (
	"leaf/chanrpc"
	"leaf/gate"
	"leaf/log"
	"encoding/binary"
	"common/proto"
)

type RoomModule interface{
	GetChanRPC() *chanrpc.Server
	GetId() int
	GetClientCount() int32
	RemovePlayer(roleid int32)
}
func HandleMsgData(args []interface{})error{
	a:= args[0].(gate.Agent)
	if a.Gate().Processor != nil {
		data := args[1].([]byte)
		var id = binary.BigEndian.Uint16(data)
		_,ok:=a.UserData().(int32)
		if id != uint16(proto.PCK_C2S_Login_ID) && !ok {
			log.Debug("没有登录发送包")
			return nil
		}
		msg, err := a.Gate().Processor.Unmarshal(data)
		if err != nil {
			log.Error("解包错误:%v",err)
		}

		err =a.Gate().Processor.Route(msg, a)
		if err != nil {
			log.Error("解包错误:%v",err)
		}
	}
	return  nil
}