package internal

import (
	"gameserver/common"
	"consts"
	"github.com/kataras/iris/core/errors"
)



var (
	RoomSN int32 =0//房间序号
	roomMap = map[int32]*RoomData{}
)
type RoomData struct {
	Id int32
	RoomType consts.ROOM_TYPE
	Module common.RoomModule
	Num  int32 						//房间人数
}

func getNewRoomId(args []interface{})  (interface{},error){
	RoomSN++
	return  RoomSN,nil
}
func getBestRoomId(args []interface{}) (interface{},error){
	t:=args[0].(consts.ROOM_TYPE)
	var id int32 =0
	for _,v :=range roomMap{
		if v.RoomType== t{
			if v.Num<consts.ROOM_MAX_NUM{
				id=v.Id
			}
		}
	}
	return  id,nil
}
func getModuleByRoomId(args[] interface{})(interface{},error){
	roomId:=args[0].(int32)
	room:=GetRoom(roomId)
	if room==nil {
		return  nil,errors.New("房间的module为空")
	}
	return  room.Module,nil;
}
func GetRoom(roomId int32) *RoomData{
	room, ok := roomMap[roomId]
	if ok  {
		return  room;
	}
	return  nil;
}
func onCreateRoom(args[] interface{})error{
	m:=args[0].(common.RoomModule)
	roomId:=args[1].(int32)
	roomType:=args[2].(consts.ROOM_TYPE)
	room:=GetRoom(roomId)
	if room !=nil{
		room.Module=m
		room.RoomType= roomType
	}else{
		room =&RoomData{Id:roomId,Module:m,RoomType:roomType}
		roomMap[roomId]= room
	}
	return  nil;
}
//人数发生变化
func onRoomNumChange(args[]interface{}) error{
	roomId:=args[0].(int32)
	num:=args[1].(int32)
	room:=GetRoom(roomId)
	if room!=nil{
		room.Num=num
	}
	return  nil
}

