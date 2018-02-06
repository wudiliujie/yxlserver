package internal

import (
	"leaf/module"
	"gameserver/base"
	"gameserver/gate"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	gate.Module.Gate.AgentChanRPC=ChanRPC
}

func (m *Module) OnDestroy() {

}
