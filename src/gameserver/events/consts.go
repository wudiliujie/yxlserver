package events

type EventType int32

const (
	Net_AgentInit    EventType = 1
	Net_AgentDestroy EventType = 2
	Net_ReceiveMsg   EventType = 3

	DB_LoadRoleData EventType = 4
)
