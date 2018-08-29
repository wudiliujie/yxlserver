package agentManage

import "leaf/gate"

var agents = make(map[int32]gate.Agent)

func AddAgent(gate gate.Agent) {
	agents[gate.GetId()] = gate
}
func RemoveAgent(gate gate.Agent) {
	delete(agents, gate.GetId())
}
