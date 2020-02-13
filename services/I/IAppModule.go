package I

type IAppModule interface {
	OnAgentInit(args []interface{})
	OnAgentDestroy(args []interface{})
	OnReceiveMsg(args []interface{})
	OnInitComplete(args []interface{})
}
