package app

import (
	"gameserver/app/internal"
)

var (
	Module  = new(internal.AppModule)
	ChanRPC = internal.ChanRPC
)
