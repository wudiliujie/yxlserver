package gate

import (
	"net"
	"leaf/chanrpc"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	Gate() *Gate
	SetChanRPC(*chanrpc.Server)
}
