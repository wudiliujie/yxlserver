package client

import (

	"common/proto"
	"leaf/log"
	"common/msg"
	"common/errmsg"
	"sync/atomic"
)
var (
	count=int32(0)

)
func init() {
	msg.Processor.SetHandler(&proto.S2C_Login{}, handleLogin)
	msg.Processor.SetHandler(&proto.S2C_GetInfo{}, handleGetInfo)
	msg.Processor.SetHandler(&proto.S2C_EnterRoom{}, handleEnterRoom)
}
func GetGID() uint64 {
	/*b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)*//*
	return n*/
	return  0
}
func handleLogin(args []interface{}) {


	recvMsg := args[0].(*proto.S2C_Login)
	agent:=args[1].(*Agent)
	//atomic.AddInt32(&TestCount1,1)
	//log.Debug("登录返回:%v",atomic.LoadInt32(&TestCount1))
	atomic.AddInt32(&LoginCount,1)
	if recvMsg.Tag !=errmsg.SYS_SUCCESS {
		atomic.AddInt32(&LoginCount_Fail,1)
		Close()
		log.Error("login is error: %v, please input login cmd", recvMsg.Tag)
		return
	}
	log.Debug("%v信息:%v-->%v--->%v--->%v--->%v",GetGID(), atomic.LoadInt32(&LoginCount),atomic.LoadInt32(&LoginCount_Fail),atomic.LoadInt32(&GetInfoCount),atomic.LoadInt32(&EnterRoomCount),atomic.LoadInt32(&EnterRoomCount_Fail))
	getinfo:=proto.C2S_GetInfo{};

	agent.WriteMsg(&getinfo);


}
func handleGetInfo(args []interface{}){
	//recvMsg := args[0].(*proto.S2C_GetInfo)
	agent:=args[1].(*Agent)
	//log.Debug("获取信息返回")
	atomic.AddInt32(&GetInfoCount,1)
	log.Debug("%v信息:%v-->%v--->%v--->%v--->%v",GetGID(), atomic.LoadInt32(&LoginCount),atomic.LoadInt32(&LoginCount_Fail),atomic.LoadInt32(&GetInfoCount),atomic.LoadInt32(&EnterRoomCount),atomic.LoadInt32(&EnterRoomCount_Fail))

	sendmsg:=&proto.C2S_EnterRoom{RoomType:0}
	agent.WriteMsg(sendmsg)
}
func handleEnterRoom(args[]interface{}){
	recvMsg := args[0].(*proto.S2C_EnterRoom)
	//agent:=args[1].(*Agent)

	atomic.AddInt32(&EnterRoomCount,1)
	if recvMsg.Tag !=errmsg.SYS_SUCCESS{
		//log.Debug("进入房间编号:%v",recvMsg.RoomId)
		log.Debug("登录房间结果%v",recvMsg.Tag)
		atomic.AddInt32(&EnterRoomCount_Fail,1)
	}

	if recvMsg.RoomId==500{
		atomic.AddInt32(&count,1);
		log.Debug( "EnterRoom:%v",atomic.LoadInt32(&count))
	}
	log.Debug("%v信息:%v-->%v--->%v--->%v--->%v",GetGID(), atomic.LoadInt32(&LoginCount),atomic.LoadInt32(&LoginCount_Fail),atomic.LoadInt32(&GetInfoCount),atomic.LoadInt32(&EnterRoomCount),atomic.LoadInt32(&EnterRoomCount_Fail))


}

