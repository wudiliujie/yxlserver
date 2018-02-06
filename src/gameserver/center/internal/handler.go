package internal

import (
	"gameserver/conf"
	"gameserver/room"
	"errors"
	"gameserver/common"
	"leaf/chanrpc"
	"leaf/gate"
	"leaf/log"
	"common/msg"
	"common/proto"
	"gameserver/players"
	"reflect"
	"gameserver/db"
	"strconv"
)

var (
	roomModuleMap = map[string]room.Module{}
	routeMoreMap = map[interface{}]interface{}{}
)

func handleRpc(id interface{}, f interface{}, fType int) {
	//cluster.SetRoute(id, ChanRPC)
	ChanRPC.RegisterFromType(id, f, fType)
}

func init() {

	handleRpc("C2S_Login",Login,chanrpc.FuncCommon)

	handleRpc(reflect.TypeOf(&proto.C2S_GetInfo{}),RolePlayerMsg,chanrpc.FuncRoute)

	//skeleton.RegisterChanRPC("CreatePlayer",handleCreatePlayer);
	handleRpc("handleMsgData",common.HandleMsgData,chanrpc.FuncCommon)
	handleRpc("NewAgent",rpcNewAgent,chanrpc.FuncCommon)
	handleRpc("CloseAgent",rpcCloseAgent,chanrpc.FuncCommon)
	//skeleton.RegisterChanRPC("PlayerLogout", PlayerLogout)
	msg.Processor.SetHandler(&proto.C2S_Login{}, handleLogin)
}
func rpcNewAgent(args []interface {}) {
	a:= args[0].(gate.Agent)
	log.Debug("连接成功%v",a.RemoteAddr())
}
func rpcCloseAgent(args []interface {}) {
	a:= args[0].(gate.Agent)
	log.Debug("断开连接%v",a.RemoteAddr())
}


func handleLogin(args []interface{}){
	log.Debug("登录消息");
	recvMsg := args[0].(*proto.C2S_Login)
	log.Debug("loing:%s>>>>%s",recvMsg.Name,recvMsg.Password)
	agent := args[1].(gate.Agent)
	accountid := 1
	agent.SetUserData(accountid)
	player:=players.GetPlayer(accountid);
	if player!=nil{ //
		//已经在线,替换agent
		player.ReplaceAgent(agent);

	}
	player= players. CreatePlayer(accountid);
	if player==nil{
		sendMsg := &proto.S2C_Login{Err:"创建player错误"}
		agent.WriteMsg(sendMsg)
	}
	player.Agent= agent;
	//读取
	skeleton.Go(func(){
		//读取mongo
		data:=db.ReadRoleInfo(1);
		player.InitData(data);
	},func(){
		EnterRoomModule(player);
	})

}
//进入大厅
func EnterRoomModule(player* players.Player){
	module := room.GetBestModule()
	if module == nil {
		sendMsg := &proto.S2C_Login{Err:"room模块错误"}
		player.Agent.WriteMsg(sendMsg)
		return
	}
	newArgs := []interface{}{module,player,func( err error){
		log.Debug("进入模块成功:%v",module.Name)
		if err==nil{
			sendMsg := &proto.S2C_Login{Err:"",AccountId:strconv.Itoa(player.AccountId)}
			(player.Agent).WriteMsg(sendMsg)

		}else {
			log.Debug(err.Error())
			sendMsg := &proto.S2C_Login{Err:err.Error()}
			player.Agent.WriteMsg(sendMsg)
		}
	}}
	skeleton.AsynCall(module.GetChanRPC(), "EnterModule", newArgs...)

}
func GetChatInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{common.GetClientCount(), conf.Server.ListenAddr}, nil
}


func EnterRoom(args []interface{}) {
	roomName := args[0].(string)

	module := roomModuleMap[roomName]
	if module == nil {
		module = room.GetBestModule()
		if module == nil {
			retFunc := args[len(args)-1].(chanrpc.ExtRetFunc)
			retFunc(nil, errors.New("get best room module rpc is fail"))
			return
		}

		roomModuleMap[roomName] = module
	}

	newArgs := []interface{}{module}
	newArgs = append(newArgs, args...)
	skeleton.AsynCall(module.GetChanRPC(), "EnterRoom", newArgs...)
}

func Login(args []interface{}) ([]interface{}, error) {


	return []interface{}{common.GetClientCount(), conf.Server.ListenAddr}, nil
}

func RouteSingle(args []interface{}) {
	id := args[0]
	roomName := args[1].(string)

	module := roomModuleMap[roomName]
	if module == nil {
		retFunc := args[len(args)-1].(chanrpc.ExtRetFunc)
		retFunc(nil, errors.New("this room is not exist"))
		return
	}

	args = append([]interface{}{module}, args[1:]...)
	skeleton.AsynCall(module.GetChanRPC(), id, args...)
}

func RouteMore(args []interface{}) {
	id := args[0]
	roomNames := args[1].([]string)
	retFunc := args[len(args)-1].(chanrpc.ExtRetFunc)

	id = routeMoreMap[id]
	for _, roomName := range roomNames {
		module := roomModuleMap[roomName]
		if module != nil {
			args = append([]interface{}{module, roomName}, args[2:]...)
			module.GetChanRPC().Go(id, args...)
		}
	}
	retFunc(nil, nil)
}
func RolePlayerMsg(args [] interface{}){
	id := args[0]
	agent := args[2].(gate.Agent)
	accountid:= agent.UserData().(int);
	player:= players.GetPlayer(accountid)
	if  player != nil{
		if player.Module!= nil{
			skeleton.AsynCall( (player.Module).GetChanRPC(), id, args...)
		}
	}
}