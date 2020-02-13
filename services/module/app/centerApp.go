package app

type CenterAppModule struct {
	BaseAppModule
}

func CreateCenterAppModule() *CenterAppModule {
	ret := &CenterAppModule{}
	ret.self = ret
	return ret
}
