package msg

import (
	"github.com/golang/protobuf/proto"
	"github.com/wudiliujie/common/network"
	"github.com/wudiliujie/common/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(PCK_C2SLogin_ID, func() network.IMessage { return new(C2SLogin) })
	Processor.Register(PCK_S2CLogin_ID, func() network.IMessage { return new(S2CLogin) })
	Processor.Register(PCK_S2CRoleInfo_ID, func() network.IMessage { return new(S2CRoleInfo) })
	Processor.Register(PCK_S2SRegisterServerInfo_ID, func() network.IMessage { return new(S2SRegisterServerInfo) })
	Processor.Register(PCK_S2SServerInfo_ID, func() network.IMessage { return new(S2SServerInfo) })
	Processor.Register(PCK_S2CKill_ID, func() network.IMessage { return new(S2CKill) })
	Processor.Register(PCK_C2SLoginEnd_ID, func() network.IMessage { return new(C2SLoginEnd) })
}

const PCK_C2SLogin_ID = 1 //登录
//登录
type C2SLogin struct {
	//登录时间(秒)
	LoginTime int64 `protobuf:"varint,1,opt,name=LoginTime,proto3" json:"LoginTime,omitempty"`
	//账号角色id
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//Sign
	Sign string `protobuf:"bytes,3,opt,name=Sign,proto3" json:"Sign,omitempty"`
}

func (m *C2SLogin) Reset()         { *m = C2SLogin{} }
func (m *C2SLogin) String() string { return proto.CompactTextString(m) }
func (*C2SLogin) ProtoMessage()    {}
func (m *C2SLogin) GetId() uint16  { return PCK_C2SLogin_ID }

const PCK_S2CLogin_ID = 2 //登录返回
//登录返回
type S2CLogin struct {
	//返回结果
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CLogin) Reset()         { *m = S2CLogin{} }
func (m *S2CLogin) String() string { return proto.CompactTextString(m) }
func (*S2CLogin) ProtoMessage()    {}
func (m *S2CLogin) GetId() uint16  { return PCK_S2CLogin_ID }

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

//用户字符串属性
type StrAttr struct {
	//Key
	K int32 `protobuf:"varint,1,opt,name=K,proto3" json:"K,omitempty"`
	//value
	V string `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *StrAttr) Reset()         { *m = StrAttr{} }
func (m *StrAttr) String() string { return proto.CompactTextString(m) }
func (*StrAttr) ProtoMessage()    {}

const PCK_S2CRoleInfo_ID = 3 //用户角色信息
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
func (m *S2CRoleInfo) GetId() uint16  { return PCK_S2CRoleInfo_ID }

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

const PCK_S2SRegisterServerInfo_ID = 4 //注册服务器信息
//注册服务器信息
type S2SRegisterServerInfo struct {
	//服务器信息
	Info *ServerInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2SRegisterServerInfo) Reset()         { *m = S2SRegisterServerInfo{} }
func (m *S2SRegisterServerInfo) String() string { return proto.CompactTextString(m) }
func (*S2SRegisterServerInfo) ProtoMessage()    {}
func (m *S2SRegisterServerInfo) GetId() uint16  { return PCK_S2SRegisterServerInfo_ID }

const PCK_S2SServerInfo_ID = 5 //发送当前链接的游戏服务器信息到网关
//发送当前链接的游戏服务器信息到网关
type S2SServerInfo struct {
	//服务器信息
	Info []*ServerInfo `protobuf:"bytes,1,rep,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2SServerInfo) Reset()         { *m = S2SServerInfo{} }
func (m *S2SServerInfo) String() string { return proto.CompactTextString(m) }
func (*S2SServerInfo) ProtoMessage()    {}
func (m *S2SServerInfo) GetId() uint16  { return PCK_S2SServerInfo_ID }

//加载用户信息回调
type D2SLoadRoleCallback struct {
	//加载结果
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//用户编号
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//登录区编号
	LoginAreaId int64 `protobuf:"varint,3,opt,name=LoginAreaId,proto3" json:"LoginAreaId,omitempty"`
	//数据库数据
	Data *RoleDbData `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *D2SLoadRoleCallback) Reset()         { *m = D2SLoadRoleCallback{} }
func (m *D2SLoadRoleCallback) String() string { return proto.CompactTextString(m) }
func (*D2SLoadRoleCallback) ProtoMessage()    {}

//用户数据库数据
type RoleDbData struct {
	//int属性
	A *IntAttr `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	//Str属性
	B *StrAttr `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (m *RoleDbData) Reset()         { *m = RoleDbData{} }
func (m *RoleDbData) String() string { return proto.CompactTextString(m) }
func (*RoleDbData) ProtoMessage()    {}

const PCK_S2CKill_ID = 6 //用户提下线
//用户提下线
type S2CKill struct {
	//Tag
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CKill) Reset()         { *m = S2CKill{} }
func (m *S2CKill) String() string { return proto.CompactTextString(m) }
func (*S2CKill) ProtoMessage()    {}
func (m *S2CKill) GetId() uint16  { return PCK_S2CKill_ID }

const PCK_C2SLoginEnd_ID = 7 //用户登录成功
//用户登录成功
type C2SLoginEnd struct {
}

func (m *C2SLoginEnd) Reset()         { *m = C2SLoginEnd{} }
func (m *C2SLoginEnd) String() string { return proto.CompactTextString(m) }
func (*C2SLoginEnd) ProtoMessage()    {}
func (m *C2SLoginEnd) GetId() uint16  { return PCK_C2SLoginEnd_ID }
