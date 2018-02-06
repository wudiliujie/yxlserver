// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	C2S_Login
	S2C_Login
	C2S_GetInfo
	S2C_GetInfo
	IntAttr
	StrAttr
	RoleDbInfo
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type PCK int32

const (
	PCK_Default_ID     PCK = 0
	PCK_C2S_Login_ID   PCK = 1
	PCK_S2C_Login_ID   PCK = 2
	PCK_C2S_GetInfo_ID PCK = 3
	PCK_S2C_GetInfo_ID PCK = 4
	PCK_MAX_ID         PCK = 100
)

var PCK_name = map[int32]string{
	0:   "Default_ID",
	1:   "C2S_Login_ID",
	2:   "S2C_Login_ID",
	3:   "C2S_GetInfo_ID",
	4:   "S2C_GetInfo_ID",
	100: "MAX_ID",
}
var PCK_value = map[string]int32{
	"Default_ID":     0,
	"C2S_Login_ID":   1,
	"S2C_Login_ID":   2,
	"C2S_GetInfo_ID": 3,
	"S2C_GetInfo_ID": 4,
	"MAX_ID":         100,
}

func (x PCK) String() string {
	return proto1.EnumName(PCK_name, int32(x))
}
func (PCK) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type C2S_Login struct {
	Name     string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,json=password" json:"Password,omitempty"`
}

func (m *C2S_Login) Reset()                    { *m = C2S_Login{} }
func (m *C2S_Login) String() string            { return proto1.CompactTextString(m) }
func (*C2S_Login) ProtoMessage()               {}
func (*C2S_Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S_Login) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *C2S_Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type S2C_Login struct {
	Err       string `protobuf:"bytes,1,opt,name=Err,json=err" json:"Err,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
}

func (m *S2C_Login) Reset()                    { *m = S2C_Login{} }
func (m *S2C_Login) String() string            { return proto1.CompactTextString(m) }
func (*S2C_Login) ProtoMessage()               {}
func (*S2C_Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C_Login) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *S2C_Login) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type C2S_GetInfo struct {
}

func (m *C2S_GetInfo) Reset()                    { *m = C2S_GetInfo{} }
func (m *C2S_GetInfo) String() string            { return proto1.CompactTextString(m) }
func (*C2S_GetInfo) ProtoMessage()               {}
func (*C2S_GetInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type S2C_GetInfo struct {
	Hp int32 `protobuf:"varint,1,opt,name=hp" json:"hp,omitempty"`
	Mp int32 `protobuf:"varint,2,opt,name=mp" json:"mp,omitempty"`
}

func (m *S2C_GetInfo) Reset()                    { *m = S2C_GetInfo{} }
func (m *S2C_GetInfo) String() string            { return proto1.CompactTextString(m) }
func (*S2C_GetInfo) ProtoMessage()               {}
func (*S2C_GetInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *S2C_GetInfo) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *S2C_GetInfo) GetMp() int32 {
	if m != nil {
		return m.Mp
	}
	return 0
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// db
type IntAttr struct {
	K int32 `protobuf:"varint,1,opt,name=k" json:"k,omitempty"`
	V int64 `protobuf:"varint,2,opt,name=v" json:"v,omitempty"`
}

func (m *IntAttr) Reset()                    { *m = IntAttr{} }
func (m *IntAttr) String() string            { return proto1.CompactTextString(m) }
func (*IntAttr) ProtoMessage()               {}
func (*IntAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IntAttr) GetK() int32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *IntAttr) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

type StrAttr struct {
	K int32  `protobuf:"varint,1,opt,name=k" json:"k,omitempty"`
	V string `protobuf:"bytes,2,opt,name=v" json:"v,omitempty"`
}

func (m *StrAttr) Reset()                    { *m = StrAttr{} }
func (m *StrAttr) String() string            { return proto1.CompactTextString(m) }
func (*StrAttr) ProtoMessage()               {}
func (*StrAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StrAttr) GetK() int32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *StrAttr) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

type RoleDbInfo struct {
	Roleid  int32      `protobuf:"varint,1,opt,name=roleid" json:"roleid,omitempty"`
	IntAttr []*IntAttr `protobuf:"bytes,2,rep,name=intAttr" json:"intAttr,omitempty"`
	StrAttr []*StrAttr `protobuf:"bytes,3,rep,name=strAttr" json:"strAttr,omitempty"`
}

func (m *RoleDbInfo) Reset()                    { *m = RoleDbInfo{} }
func (m *RoleDbInfo) String() string            { return proto1.CompactTextString(m) }
func (*RoleDbInfo) ProtoMessage()               {}
func (*RoleDbInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RoleDbInfo) GetRoleid() int32 {
	if m != nil {
		return m.Roleid
	}
	return 0
}

func (m *RoleDbInfo) GetIntAttr() []*IntAttr {
	if m != nil {
		return m.IntAttr
	}
	return nil
}

func (m *RoleDbInfo) GetStrAttr() []*StrAttr {
	if m != nil {
		return m.StrAttr
	}
	return nil
}

func init() {
	proto1.RegisterType((*C2S_Login)(nil), "proto.C2S_Login")
	proto1.RegisterType((*S2C_Login)(nil), "proto.S2C_Login")
	proto1.RegisterType((*C2S_GetInfo)(nil), "proto.C2S_GetInfo")
	proto1.RegisterType((*S2C_GetInfo)(nil), "proto.S2C_GetInfo")
	proto1.RegisterType((*IntAttr)(nil), "proto.IntAttr")
	proto1.RegisterType((*StrAttr)(nil), "proto.StrAttr")
	proto1.RegisterType((*RoleDbInfo)(nil), "proto.RoleDbInfo")
	proto1.RegisterEnum("proto.PCK", PCK_name, PCK_value)
}

func init() { proto1.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcb, 0x6b, 0xc2, 0x40,
	0x10, 0xc6, 0x9b, 0xac, 0xaf, 0x8c, 0x36, 0xc8, 0x1c, 0x4a, 0x28, 0x3d, 0x48, 0xa0, 0x20, 0x85,
	0x7a, 0xb0, 0xc7, 0x9e, 0x44, 0x4b, 0x09, 0x7d, 0x20, 0xeb, 0xa5, 0x37, 0x89, 0xba, 0x3e, 0xd0,
	0x64, 0x97, 0xcd, 0x6a, 0xfb, 0xe7, 0x97, 0xd9, 0x6c, 0xc4, 0x5e, 0x7a, 0x4a, 0xbe, 0x6f, 0x7e,
	0xf3, 0xe0, 0x5b, 0x08, 0xb2, 0x62, 0x33, 0x50, 0x5a, 0x1a, 0x89, 0x75, 0xfb, 0x89, 0x9f, 0x21,
	0x18, 0x0f, 0x67, 0xf3, 0x77, 0xb9, 0xd9, 0xe5, 0x88, 0x50, 0xfb, 0x4c, 0x33, 0x11, 0x79, 0x3d,
	0xaf, 0x1f, 0xf0, 0x5a, 0x9e, 0x66, 0x02, 0x6f, 0xa1, 0x35, 0x4d, 0x8b, 0xe2, 0x5b, 0xea, 0x55,
	0xe4, 0x5b, 0xbf, 0xa5, 0x9c, 0xa6, 0xe6, 0xd9, 0x70, 0xec, 0x9a, 0xbb, 0xc0, 0x5e, 0xb4, 0x76,
	0xbd, 0x4c, 0x68, 0x8d, 0x77, 0x10, 0x8c, 0x96, 0x4b, 0x79, 0xcc, 0x4d, 0x52, 0xf5, 0x06, 0x69,
	0x65, 0xc4, 0xd7, 0xd0, 0xa6, 0xcd, 0xaf, 0xc2, 0x24, 0xf9, 0x5a, 0xc6, 0x8f, 0xd0, 0xa6, 0x59,
	0x4e, 0x62, 0x08, 0xfe, 0x56, 0xd9, 0x61, 0x75, 0xee, 0x6f, 0x15, 0xe9, 0x4c, 0xd9, 0x21, 0x75,
	0xee, 0x67, 0x2a, 0xbe, 0x87, 0x66, 0x92, 0x9b, 0x91, 0x31, 0x1a, 0x3b, 0xe0, 0xed, 0x1d, 0xe9,
	0xed, 0x49, 0x9d, 0x2c, 0xc7, 0xb8, 0x77, 0x22, 0x6c, 0x66, 0xf4, 0x7f, 0x58, 0x40, 0xd8, 0x0f,
	0x00, 0x97, 0x07, 0x31, 0x59, 0xd8, 0xdd, 0x37, 0xd0, 0xd0, 0xf2, 0x20, 0x76, 0x2b, 0x87, 0x3b,
	0x85, 0x7d, 0x68, 0xee, 0xca, 0x9d, 0x91, 0xdf, 0x63, 0xfd, 0xf6, 0x30, 0x2c, 0xb3, 0x1c, 0xb8,
	0x4b, 0x78, 0x55, 0x26, 0xb2, 0x28, 0xd7, 0x46, 0xec, 0x0f, 0xe9, 0x8e, 0xe1, 0x55, 0xf9, 0x21,
	0x03, 0x36, 0x1d, 0xbf, 0x61, 0x08, 0x30, 0x11, 0xeb, 0xf4, 0x78, 0x30, 0xf3, 0x64, 0xd2, 0xbd,
	0xc2, 0x2e, 0x74, 0xce, 0xcf, 0x42, 0x8e, 0x47, 0xce, 0x39, 0x6b, 0x72, 0x7c, 0x44, 0x08, 0x2f,
	0x02, 0x24, 0x8f, 0x91, 0x77, 0x91, 0x22, 0x79, 0x35, 0x04, 0x68, 0x7c, 0x8c, 0xbe, 0xe8, 0x7f,
	0xb5, 0x68, 0xd8, 0x33, 0x9e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xdb, 0x09, 0x25, 0x09,
	0x02, 0x00, 0x00,
}
