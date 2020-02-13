package msg

import (
	"github.com/golang/protobuf/proto"
	"github.com/wudiliujie/common/network"
	"github.com/wudiliujie/common/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(PCKID_C2SLogin, func() network.IMessage { return new(C2SLogin) })
	Processor.Register(PCKID_S2CLogin, func() network.IMessage { return new(S2CLogin) })
	Processor.Register(PCKID_S2CRoleInfo, func() network.IMessage { return new(S2CRoleInfo) })
	Processor.Register(PCKID_S2SRegisterServerInfo, func() network.IMessage { return new(S2SRegisterServerInfo) })
	Processor.Register(PCKID_S2SServerInfo, func() network.IMessage { return new(S2SServerInfo) })
}

const PCKID_C2SLogin = 1 //登录
//登录
type C2SLogin struct {
	//账号id
	AccountId int64 `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	//账号角色id
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//token
	Token string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (m *C2SLogin) Reset()         { *m = C2SLogin{} }
func (m *C2SLogin) String() string { return proto.CompactTextString(m) }
func (*C2SLogin) ProtoMessage()    {}
func (m *C2SLogin) GetId() uint16  { return PCKID_C2SLogin }

const PCKID_S2CLogin = 2 //登录返回
//登录返回
type S2CLogin struct {
	//返回结果
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CLogin) Reset()         { *m = S2CLogin{} }
func (m *S2CLogin) String() string { return proto.CompactTextString(m) }
func (*S2CLogin) ProtoMessage()    {}
func (m *S2CLogin) GetId() uint16  { return PCKID_S2CLogin }

//用户int属性
type IntAttr struct {
	//Key
	K int32 `protobuf:"varint,1,opt,name=K,proto3" json:"K,omitempty"`
	//value
	V int64 `protobuf:"varint,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *IntAttr) Reset()         { *m = IntAttr{} }
func (m *IntAttr) String() string { return proto.CompactTextString(m) }
func (*IntAttr) ProtoMessage()    {}

const PCKID_S2CRoleInfo = 3 //用户角色信息
//用户角色信息
type S2CRoleInfo struct {
	//用户编号
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//用户属性
	A []*IntAttr `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
}

func (m *S2CRoleInfo) Reset()         { *m = S2CRoleInfo{} }
func (m *S2CRoleInfo) String() string { return proto.CompactTextString(m) }
func (*S2CRoleInfo) ProtoMessage()    {}
func (m *S2CRoleInfo) GetId() uint16  { return PCKID_S2CRoleInfo }

//服务器信息
type ServerInfo struct {
	//服务器编号
	ServerId int64 `protobuf:"varint,1,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	//服务器类型
	ServerType int64 `protobuf:"varint,2,opt,name=ServerType,proto3" json:"ServerType,omitempty"`
	//服务器名字
	ServerName string `protobuf:"bytes,3,opt,name=ServerName,proto3" json:"ServerName,omitempty"`
	//服务器地址
	ServerAddr string `protobuf:"bytes,4,opt,name=ServerAddr,proto3" json:"ServerAddr,omitempty"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}

const PCKID_S2SRegisterServerInfo = 4 //注册服务器信息
//注册服务器信息
type S2SRegisterServerInfo struct {
	//服务器信息
	Info *ServerInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2SRegisterServerInfo) Reset()         { *m = S2SRegisterServerInfo{} }
func (m *S2SRegisterServerInfo) String() string { return proto.CompactTextString(m) }
func (*S2SRegisterServerInfo) ProtoMessage()    {}
func (m *S2SRegisterServerInfo) GetId() uint16  { return PCKID_S2SRegisterServerInfo }

const PCKID_S2SServerInfo = 5 //发送当前链接的游戏服务器信息到网关
//发送当前链接的游戏服务器信息到网关
type S2SServerInfo struct {
	//服务器信息
	Info []*ServerInfo `protobuf:"bytes,1,rep,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2SServerInfo) Reset()         { *m = S2SServerInfo{} }
func (m *S2SServerInfo) String() string { return proto.CompactTextString(m) }
func (*S2SServerInfo) ProtoMessage()    {}
func (m *S2SServerInfo) GetId() uint16  { return PCKID_S2SServerInfo }
