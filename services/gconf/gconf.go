package gconf

import (
	"github.com/wudiliujie/common/network"
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000

	PendingWriteNum        = 3000
	MaxMsgLen       uint32 = 409600
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false

	// agent conf
	AgentGoLen              = 500
	AgentTimerDispatcherLen = 50
	AgentAsyncCallLen       = 50
	AgentChanRPCLen         = 50

	// skeleton conf

	Processor network.IProcessor
)
