package main

import (
	"common"
	"gameserver/conf"
	"leaf"
	lconf "leaf/conf"
	"leaf/log"
	"leaf/module"
	"os"

	"gameserver/app"
	_ "gameserver/common/actions/login"
	"gameserver/db"
	"gameserver/gate"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatal("os args of len(%v) less than 2", argsLen)
	}

	confFileName := os.Args[1]
	conf.Init(confFileName)

	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	lconf.ServerName = conf.Server.ServerName
	lconf.ListenAddr = conf.Server.ListenAddr
	lconf.ConnAddrs = conf.Server.ConnAddrs
	lconf.PendingWriteNum = conf.Server.PendingWriteNum
	lconf.HeartBeatInterval = conf.HeartBeatInterval

	common.Init()

	modules := []module.Module{app.Module, gate.Module, db.Module}
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	leaf.Run(modules...)
}
