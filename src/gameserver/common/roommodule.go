package common

import (
	"leaf/chanrpc"
	"leaf/gate"
)

type RoomModule interface{
	GetChanRPC() *chanrpc.Server
}
func HandleMsgData(args []interface{}) error{
	a:= args[0].(gate.Agent)
	if a.Gate().Processor != nil {
		data := args[1].([]byte)
		msg, err := a.Gate().Processor.Unmarshal(data)
		if err != nil {
			return err
		}

		err =a.Gate().Processor.Route(msg, a)
		if err != nil {
			return err
		}
	}
	return nil
}