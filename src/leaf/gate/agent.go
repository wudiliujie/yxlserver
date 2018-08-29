package gate

import (
	"net"
)

type Agent interface {
	GetId() int32
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	Gate() *Gate
}
