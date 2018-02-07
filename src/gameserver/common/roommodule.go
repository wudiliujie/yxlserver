package common

import (
	"leaf/chanrpc"
	"leaf/gate"
	"leaf/log"
)

type RoomModule interface{
	GetChanRPC() *chanrpc.Server
	GetId() int
	GetClientCount() int32
}
func HandleMsgData(args []interface{})error{
	a:= args[0].(gate.Agent)
	if a.Gate().Processor != nil {
		data := args[1].([]byte)
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