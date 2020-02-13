package conf

import "yxlserver/services/consts"

var Server struct {
	ServerId   int64
	ServerType consts.ServerType
	ServerName string
	WSAddr     string
	MaxConnNum int
}

func init() {
	Server.WSAddr = "0.0.0.0:7777"
	Server.MaxConnNum = 1000
}
