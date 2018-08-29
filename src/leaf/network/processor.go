package network

type Processor interface {
	// must goroutine safe
	Route(msg interface{}, userData interface{}) error
	//获取消息编号
	GetMsgId(msg interface{}) int32
	// must goroutine safe
	Unmarshal(data []byte) (interface{}, error)
	// must goroutine safe
	Marshal(msg interface{}) ([][]byte, error)
}
