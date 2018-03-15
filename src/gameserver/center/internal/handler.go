package internal

import (

	"gameserver/common"
	"gameserver/conf"
	"consts"
	"math"
	"github.com/kataras/iris/core/errors"
)

var (
	routeMoreMap = map[interface{}]interface{}{}
)

func handleRpc(id interface{}, f interface{}, fType int) {
	//cluster.SetRoute(id, ChanRPC)
	ChanRPC.Register(id, f)
}

func init() {
	ChanRPC.Register("handleMsgData", common.HandleMsgData)
	ChanRPC.Register("NewAgent", rpcNewAgent)
	ChanRPC.Register("CloseAgent", rpcCloseAgent)
	ChanRPC.Register(consts.Center_Rpc_GetBestRoomId, getBestRoomId)
	ChanRPC.Register(consts.Center_Rpc_GetModuleByRoomId, getModuleByRoomId)
	ChanRPC.Register(consts.Center_Rpc_GetNewRoomId, getNewRoomId)
	ChanRPC.Register(consts.Center_Rpc_OnCreateRoom, onCreateRoom)
	ChanRPC.Register(consts.Center_Rpc_OnRoomNumChange, onRoomNumChange)
	ChanRPC.Register(consts.Center_Rpc_OnPlayerLogout, onPlayerLogout)
	ChanRPC.Register(consts.Center_Rpc_GetBestModule, GetBestModule)



}
func rpcNewAgent(args []interface{}) {
	//a := args[0].(gate.Agent)
	//log.Debug("连接成功%v-->%v", a.RemoteAddr(),atomic.LoadInt32(&TestCount))
}
func rpcCloseAgent(args []interface{}) error {
	//a := args[0].(gate.Agent)
	//log.Debug("断开连接%v", a.RemoteAddr())
	return  nil
}


func GetChatInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{common.GetClientCount(), conf.Server.ListenAddr}, nil
}
func onPlayerLogout(args [] interface{})(error){
	return  nil
}
func  GetBestModule(args [] interface{}) (interface{} ,error) {
	var minCount int32 = math.MaxInt32
	var m common.RoomModule
	var find bool =false
	for _, _module := range roomModule {
		count := _module.GetClientCount()
		if count < minCount {
			m = _module
			minCount = count
			find=true
		}
	}
	if find{
		return m,nil
	}else{
		return  nil,errors.New("没有找到Module")
	}

}

