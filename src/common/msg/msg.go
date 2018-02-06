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
}
