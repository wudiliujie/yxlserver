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
	"leaf/gate"
	"gameserver/db"
	"github.com/kataras/iris/core/errors"
	"common/proto"
	"common/errmsg"
	"leaf/network"
	"common/msg"

)

func NewModule(id int) *Module {
	skeleton := base.NewSkeleton()
	module := &Module{Skeleton: skeleton, ChanRPC: skeleton.ChanRPCServer}
	module.Name = fmt.Sprintf("module%v", id)
	module.roomMap = map[int32]*RoomData{}
	module.playerMap =map[int32]*Player{}
	module.Id=id;
	module.Processor=msg.Processor
	RegisterHandler(module.ChanRPC,module)
	return module
}

type HandleArgs struct{
	M *Module
	P * Player
}
type Module struct {
	*module.Skeleton
	ChanRPC *chanrpc.Server
	Id int
	Name          string
	clientCount   int32
	roomMap map[int32]* RoomData
	playerMap  map[int32]*Player
	Processor  network.Processor
}
func (m* Module) GetId() int{
	return m.Id
}
func (m *Module) OnInit() {
	m.Skeleton.AfterFunc(time.Duration(conf.DestroyRoomInterval/10), m.checkDestroyRoom)
}

func (m *Module) OnDestroy() {

}

func (m *Module) checkDestroyRoom() {
	/*destroyRooms := []string{}
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
	m.Skeleton.AfterFunc(time.Duration(conf.DestroyRoomInterval/10), m.checkDestroyRoom)*/
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

func (m *Module)HandleMsgData(args []interface{})error {
	log.Debug("%v处理消息",m.Name)
	a:= args[0].(gate.Agent)
	if a.Gate().Processor != nil {
		data := args[1].([]byte)
		msg, err := a.Gate().Processor.Unmarshal(data)
		if err != nil {
			log.Error("解析包错误:%v",err)
		}
		roleid:= a.UserData().(int32)
		player:= m.GetPlayer(roleid)
		args:=&HandleArgs{M:m,P:player}
		if(player!=nil){
			err =a.Gate().Processor.Route(msg, args)
			if err != nil {
				log.Error("解析包错误1:%v",err)
			}
		}
	}
	return  nil
}



func (m *Module) CloseAgent(args []interface{})error{
	agent := args[0].(gate.Agent)
	roleId:= agent.UserData().(int32);
	player:=m.GetPlayer(roleId);
	if player!=nil{
		data:= player.GetSaveData()
		m.Skeleton.Go(func(){
			db.SaveRoleInfo(roleId,data)
		}, func() {
			log.Debug("保存成功");
		});
	}
	m.AddClientCount(-1)
	return  nil;
}
//登录模块
func (m* Module) HandleLoginModule(args[] interface{}) error{
	roleid:=args[0].(int32)
	agent := args[1].(gate.Agent)

	player:= m.GetPlayer(roleid)
	if player!=nil {
		return  errors.New("帐号重复登录");
	}
	player = m.CreatePlayer(roleid,agent);
	if player==nil{
		return  errors.New("帐号重复登录1");
	}

	agent.SetChanRPC(m.ChanRPC)

	var data []byte;
	m.Skeleton.Go(func() {
		  db.ReadRoleInfo(roleid,&data)
	},func(){
		log.Debug("用户读取数据完成")
		player.InitData(&data);
		//发送登录成功
		sendmsg:=&proto.S2C_Login{ Tag:errmsg.SYS_SUCCESS}
		player.SendMsg(sendmsg)
	})

	return  nil
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

