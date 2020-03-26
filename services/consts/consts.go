package consts

type ServerType int32

const (
	ServerType_Center ServerType = 1
	ServerType_Gate   ServerType = 2
	ServerType_Server ServerType = 3
)

//区状态
type ServerState int32

const (
	ServerState_Close    ServerState = 0 //关闭
	ServerState_Starting ServerState = 1 //启动中
	ServerState_Run      ServerState = 2 //运行中
)

type PlayerState int32

const (
	PlayerState_Initing = 1
	PlayerState_OffLine = 2
	PlayerState_OnLine  = 3
	PlayerState_Error   = 4
)

type GradeType int32

const (
	GradeType_Role    GradeType = 0 //人物
	GradeType_Pet     GradeType = 1 //宠物
	GradeType_PetA    GradeType = 2 //仙侣
	GradeType_Beauty  GradeType = 3 //玄女
	GradeType_Kid     GradeType = 4 //仙童
	GradeType_SkyGod  GradeType = 5 //天神
	GradeType_Therion GradeType = 6 //圣兽
)
