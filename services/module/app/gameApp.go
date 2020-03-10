package app

type GameAppModule struct {
	BaseAppModule
}

func CreateGameAppModule() *GameAppModule {
	ret := &GameAppModule{}
	ret.self = ret
	return ret
}
