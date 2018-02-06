package internal

import (
	"leaf/module"
	"gameserver/base"
	"leaf/chanrpc"
	"gopkg.in/mgo.v2/bson"
	"leaf/log"
	"time"
	"gameserver/conf"

	"sync/atomic"
	"fmt"
	"gameserver/players"
	"leaf/gate"
	"gameserver/db"
)

func NewModule(id int) *Module {
	skeleton := base.NewSkeleton()
	module := &Module{Skeleton: skeleton, ChanRPC: skeleton.ChanRPCServer}
	module.Name = fmt.Sprintf("module%v", id)
	module.roomInfoMap = map[string]*RoomInfo{}
	module.playerMap = map[int]* players.Player{}
	RegisterHandler(module.ChanRPC,module)
	return module
}

type Module struct {
	*module.Skeleton
	ChanRPC *chanrpc.Server

	Name          string
	clientCount   int32
	roomInfoMap   map[string]*RoomInfo
	playerMap     map[int]* players.Player
}

func (m *Module) OnInit() {
	m.Skeleton.AfterFunc(time.Duration(conf.DestroyRoomInterval/10), m.checkDestroyRoom)
}

func (m *Module) OnDestroy() {

}

func (m *Module) checkDestroyRoom() {
	destroyRooms := []string{}
	nowTime := time.Now().Unix()
	for roomName, roomInfo := range m.roomInfoMap {
		if roomInfo.CheckDestroy(nowTime) {
			destroyRooms = append(destroyRooms, roomName)
		}
	}

	if len(destroyRooms) > 0 {
		for _, roomName := range destroyRooms {
			delete(m.roomInfoMap, roomName)
		}
		//cluster.Go("world", "DestroyRoom", destroyRooms)
		log.Debug("%v rooms is destroy in %v", destroyRooms, m.Name)
	}
	m.Skeleton.AfterFunc(time.Duration(conf.DestroyRoomInterval/10), m.checkDestroyRoom)
}

func (m *Module) GetChanRPC() *chanrpc.Server {
	return m.ChanRPC
}

func (m *Module) GetClientCount() int32 {
	return atomic.LoadInt32(&m.clientCount)
}

func (m *Module) AddClientCount(delta int32) {
	atomic.AddInt32(&m.clientCount, delta)
	//common.AddClientCount(delta)
}

func (m *Module) GetRoomInfo(roomName string) *RoomInfo {
	return m.roomInfoMap[roomName]
}

func (m *Module) NewRoom(name string) *RoomInfo {
	room := &RoomInfo{name: name}
	room.module = m
	room.userServerMap = map[bson.ObjectId]string{}
	m.roomInfoMap[name] = room
	return room
}

func (m *Module) PlayerEnter(player *players.Player) {
	m.playerMap[1]= player;
	player.Agent.SetChanRPC(m.ChanRPC)
}
func (m *Module) GetPlayer(roleId int)* players.Player {
	return m.playerMap[roleId];
}

func (m *Module) CloseAgent(args []interface{})(error){
	agent := args[0].(gate.Agent)
	roleId:= agent.UserData().(int);
	player:=m.GetPlayer(roleId);
	if player!=nil{
		data:= player.GetSaveData()
		m.Skeleton.Go(func(){
			db.SaveRoleInfo(roleId,data)
		}, func() {
			log.Debug("保存成功");
		});
	}
	return  nil;
}


type RoomInfo struct {
	module *Module

	name          string
	startNullTime int64
	userServerMap map[bson.ObjectId]string
}

func (r *RoomInfo) CheckDestroy(curTime int64) bool {
	if r.GetUserCount() < 1 {
		if r.startNullTime > 0 {
			if curTime-r.startNullTime >= int64(conf.DestroyRoomInterval) {
				return true
			}
		} else {
			r.startNullTime = time.Now().Unix()
		}
	} else {
		r.startNullTime = 0
	}
	return false
}

func (r *RoomInfo) GetUserCount() int {
	return len(r.userServerMap)
}

func (r *RoomInfo) IsInRoom(userId bson.ObjectId) bool {
	_, ok := r.userServerMap[userId]
	return ok
}

func (r *RoomInfo) EnterRoom(userId bson.ObjectId, serverName string) bool {
	ok := !r.IsInRoom(userId)
	if ok {
		r.userServerMap[userId] = serverName
		r.module.AddClientCount(1)
		log.Debug("%v user enter %v room, %v count user in %v", userId, r.name, r.GetUserCount(), r.module.Name)
	}
	return ok
}

func (r *RoomInfo) LeaveRoom(userId bson.ObjectId) bool {
	ok := r.IsInRoom(userId)
	if ok {
		delete(r.userServerMap, userId)
		r.module.AddClientCount(-1)
		log.Debug("%v user leave %v room, %v count user in %v", userId, r.name, r.GetUserCount(), r.module.Name)
	}
	return ok
}

