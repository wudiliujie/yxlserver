package consts

type ServerType int32

const (
	ServerType_Center ServerType = 1
	ServerType_Gate   ServerType = 2
	ServerType_Server ServerType = 3
)

//区状态
type AreaState int32

const (
	AreaState_Close    AreaState = 1 //关闭
	AreaState_Starting AreaState = 2 //启动中
	AreaState_Run      AreaState = 3 //运行中
	AreaState_Closeing AreaState = 4 //关闭中
)
