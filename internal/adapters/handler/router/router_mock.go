package router

type MockedAdapter struct {
	InitializeRoutesFn func()
	ServeFn            func(listenAdr string) error
}

func (mA *MockedAdapter) InitializeRoutes() {
	return
}

func (mA *MockedAdapter) Serve(listenAdr string) error {
	return mA.ServeFn(listenAdr)
}
