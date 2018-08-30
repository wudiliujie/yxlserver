package db

import (
	"gameserver/db/internal"
)

var (
	Module  = new(internal.DbModule)
	ChanRPC = internal.ChanRPC
)
