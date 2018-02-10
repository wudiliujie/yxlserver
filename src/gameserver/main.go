package main

import (
	"common"
	"leaf"
	lconf "leaf/conf"
	"gameserver/conf"
	"os"
	"leaf/log"
	"gameserver/center"
	"gameserver/room"
	"leaf/module"
	"gameserver/gate"
	"math/rand"
	"time"
	"gameserver/db"
	 _ "net/http/pprof"
	"net/http"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatal("os args of len(%v) less than 2", argsLen)
	}
	db.Init()
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

	modules := []module.Module{center.Module}
	modules = append(modules, room.CreateModules()...)
	modules =append(modules,gate.Module)
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	leaf.Run(modules...)
}
