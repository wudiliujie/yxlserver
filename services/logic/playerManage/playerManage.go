package playerManage

import (
	"github.com/wudiliujie/common/eventhub"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"github.com/wudiliujie/common/util"
	"yxlserver/services/I"
	"yxlserver/services/I/eventcode"
	"yxlserver/services/consts"
	"yxlserver/services/errmsg"
	"yxlserver/services/msg"
)

//一般管理的包，最好不要在调用外部的包
var playerMap = make(map[int64]I.IPlayer)

//用户数据异步加载时，保存回调函数，加载完成后清除
var playerAsyncFunc = make(map[int64][]*dataFunc)

//正在加载的用户编号
var playerLoadingMap = make(map[int64]I.IPlayer)

type dataFunc struct {
	Cb    func(I.IPlayer, ...interface{})
	Param []interface{}
}

var CreatePlayerFunc func(userId int64) I.IPlayer

//玩家管理
func Init() {
	eventhub.Subscribe(eventcode.Area_Close, OnAreaClose)
}
func OnAreaClose(args ...interface{}) {
	serverId := args[0].(int32)
	for _, player := range playerMap {
		if player.GetInt32Value(consts.AttrType_LoginAreaId) == serverId {
			player.OnAreaClose()
		}
	}
}
func GetPlayerById(userId int64) I.IPlayer {
	return playerMap[userId]
}

func SignPlayerAsyncFunc(userId int64, cb func(targetPlayer I.IPlayer, args ...interface{}), args ...interface{}) {
	if userId == 0 {
		log.Error("SignPlayerAsyncFunc userId 为0")
		return
	}
	player := GetPlayerById(userId)
	if player != nil {
		cb(player, args...)
		return
	}
	if playerAsyncFunc[userId] == nil {
		playerAsyncFunc[userId] = make([]*dataFunc, 0)
	}
	playerAsyncFunc[userId] = append(playerAsyncFunc[userId], &dataFunc{
		Cb:    cb,
		Param: args,
	})
}
func updatePlayerAsyncFunc() {
	for userId, dfs := range playerAsyncFunc {
		player := GetPlayerById(userId)
		//如果玩家数据已经加载完成则执行回调并删除回调
		if util.IsNil(player) == false && player.GetState() != consts.PlayerState_Initing {
			delete(playerAsyncFunc, userId)
			//cLog.Counter("加载玩家 %v 数据完成",id)
			for _, df := range dfs {

				df.Cb(player, df.Param...)
			}
		} else if player == nil {
			if util.IsNil(playerLoadingMap[userId]) == false {
				continue
			}
			//异步加载玩家数据并设置玩家状态为加载中
			player := CreatePlayerFunc(userId)
			log.Debug("new player:%v", userId)
			playerLoadingMap[userId] = player
			player.SetState(consts.PlayerState_Initing)
			//cLog.Counter("开始加载玩家 %v 数据",id)
			module.OnChanRpcEvent(eventcode.DB_LoadRoleData, userId)
		}
	}
}

func LoadUserCallBack(args []interface{}) {
	callBack := args[0].(*msg.D2SLoadRoleCallback)
	_player := playerLoadingMap[callBack.UserId]
	if util.IsNil(_player) == false {
		delete(playerLoadingMap, callBack.UserId)
		if callBack.Tag == errmsg.SYS_成功 {
			playerMap[_player.GetId()] = _player
			_player.InitUserData(callBack.Data)
			_player.SetIntValue(consts.AttrType_LoginAreaId, callBack.LoginAreaId)
			_player.SetState(consts.PlayerState_OffLine)
		} else {
			//todo 可能登录包都没有反应了，不应该这样做
			delete(playerAsyncFunc, _player.GetId())
		}
	}
}
func Update() {
	updatePlayerAsyncFunc()
}
