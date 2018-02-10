package msg

import (
	"leaf/network/protobuf"
	"common/proto"
)

var (
	Processor = protobuf.NewProcessor(int( proto.PCK_MAX_ID))
)

func init() {
	Processor.Register(uint16(proto.PCK_C2S_Login_ID), &proto.C2S_Login{})
	Processor.Register(uint16(proto.PCK_S2C_Login_ID) ,&proto.S2C_Login{})
	Processor.Register(uint16(proto.PCK_C2S_GetInfo_ID),&proto.C2S_GetInfo{})
	Processor.Register(uint16(proto.PCK_S2C_GetInfo_ID),&proto.S2C_GetInfo{})
	Processor.Register(uint16(proto.PCK_C2S_EnterRoom_ID),&proto.C2S_EnterRoom{})
	Processor.Register(uint16(proto.PCK_S2C_EnterRoom_ID),&proto.S2C_EnterRoom{})
	Processor.Register(uint16(proto.PCK_C2S_Fire_ID),&proto.C2S_Fire{})
	Processor.Register(uint16(proto.PCK_S2C_Fire_ID),&proto.S2C_Fire{})
}
