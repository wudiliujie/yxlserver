package serverManage

import (
	"github.com/wudiliujie/common/gate"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/network"
	"yxlserver/services/conf"
	"yxlserver/services/consts"
	"yxlserver/services/logic/actionManage"
	"yxlserver/services/msg"
)

type ServerData struct {
	agent    gate.Agent
	ServerId int64
	//服务器类型
	ServerType consts.ServerType
	//服务器名字
	ServerName string
	//服务器地址
	ServerAddr string
}

func (s *ServerData) SendPck(sendMsg network.IMessage) {
	s.agent.WriteMsg(0, sendMsg)
}

var mapServer = make(map[int64]*ServerData)
var netEvent *gate.NetEvent

const CenterServerId = 1 //
func init() {
	actionManage.RegisterAgentPacket(msg.PCKID_S2SRegisterServerInfo, OnS2SRegisterServerInfo)
}
func SetNetEvent(event *gate.NetEvent) {
	netEvent = event
}
func ConnectCenterServer() {
	if netEvent != nil {
		gate.Connect("0.0.0.0:8888", netEvent, CenterServerId)
	} else {
		log.Error("netEvent 为空")
	}
}

//当中心服务器链接
func OnServerConnect(agent gate.Agent) {
	switch agent.GetTag() {
	case CenterServerId:
		OnServerConnectCenter(agent)
		break
	default:
		//
		log.Error("未知的服务器类型：%v", agent.GetTag())
		break
	}
	//通知中心服务器，当前服务器类型
}

//上报自己信息
func OnServerConnectCenter(agent gate.Agent) {
	sendMsg := &msg.S2SRegisterServerInfo{}
	sendMsg.Info = &msg.ServerInfo{}
	sendMsg.Info.ServerId = conf.Server.ServerId
	sendMsg.Info.ServerType = int64(conf.Server.ServerType)
	sendMsg.Info.ServerName = conf.Server.ServerName
	sendMsg.Info.ServerAddr = conf.Server.WSAddr
	agent.WriteMsg(0, sendMsg)
}
func GetServerAgent(serverId int64) *ServerData {
	return mapServer[serverId]
}
func OnS2SRegisterServerInfo(agent gate.Agent, pckMsg interface{}) {
	recvMsg := pckMsg.(*msg.S2SRegisterServerInfo)
	serverData := GetServerAgent(recvMsg.Info.ServerId)
	if serverData != nil {
		//已经启动，断开新的链接，每个服务器不能重复启动
		agent.Close()
		log.Error("重复启动服务器：%v", recvMsg.Info)
		return
	}
	serverData = &ServerData{}
	serverData.agent = agent
	serverData.ServerId = recvMsg.Info.ServerId
	serverData.ServerName = recvMsg.Info.ServerName
	serverData.ServerAddr = recvMsg.Info.ServerAddr
	serverData.ServerType = consts.ServerType(recvMsg.Info.ServerType)
	//设置这个链接的服务器编号
	agent.SetTag(recvMsg.Info.ServerId)
	mapServer[recvMsg.Info.ServerId] = serverData
	//如果是游戏服务器链接成功，通知所有网关
	//如果是网关链接成功，通知网关当前所链接的所有游戏服务器信息
	if conf.Server.ServerType == consts.ServerType_Center {
		noticeGate(serverData)
	}
}

//通知网关
func noticeGate(serverData *ServerData) {
	sendMsg := &msg.S2SServerInfo{}
	if serverData.ServerType == consts.ServerType_Gate {
		for _, v := range mapServer {
			if v.ServerType == consts.ServerType_Server {
				info := &msg.ServerInfo{}
				info.ServerId = v.ServerId
				info.ServerType = int64(v.ServerType)
				info.ServerAddr = v.ServerAddr
				info.ServerName = v.ServerName
				sendMsg.Info = append(sendMsg.Info, info)
			}
		}
		if len(sendMsg.Info) > 0 {
			serverData.SendPck(sendMsg)
		}
	}
	if serverData.ServerType == consts.ServerType_Server {
		info := &msg.ServerInfo{}
		info.ServerId = serverData.ServerId
		info.ServerType = int64(serverData.ServerType)
		info.ServerAddr = serverData.ServerAddr
		info.ServerName = serverData.ServerName
		sendMsg.Info = append(sendMsg.Info, info)
		for _, v := range mapServer {
			if v.ServerType == consts.ServerType_Gate {
				v.SendPck(sendMsg)
			}
		}
	}
}
