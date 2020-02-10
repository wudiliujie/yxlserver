package msg

import "github.com/golang/protobuf/proto"

type PCK int32

const PCKID_C2SLogin PCK = 1 //登录
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
func (m *C2SLogin) GetId() int32   { return int32(PCKID_C2SLogin) }

const PCKID_S2CLogin PCK = 2 //登录返回
//登录返回
type S2CLogin struct {
	//返回结果
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CLogin) Reset()         { *m = S2CLogin{} }
func (m *S2CLogin) String() string { return proto.CompactTextString(m) }
func (*S2CLogin) ProtoMessage()    {}
func (m *S2CLogin) GetId() int32   { return int32(PCKID_S2CLogin) }

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

const PCKID_S2CRoleInfo PCK = 3 //用户角色信息
//用户角色信息
type S2CRoleInfo struct {
	//用户编号
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//用户属性
	A []IntAttr `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
}

func (m *S2CRoleInfo) Reset()         { *m = S2CRoleInfo{} }
func (m *S2CRoleInfo) String() string { return proto.CompactTextString(m) }
func (*S2CRoleInfo) ProtoMessage()    {}
func (m *S2CRoleInfo) GetId() int32   { return int32(PCKID_S2CRoleInfo) }
