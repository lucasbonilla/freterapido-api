package postgres

type MockedAdapter struct {
	InitConnFn func() error
	CloseFn    func() error
	PingFn     func() error
}

func (mA *MockedAdapter) InitConn() error {
	return mA.InitConnFn()
}

func (mA *MockedAdapter) Close() error {
	return mA.CloseFn()
}

func (mA *MockedAdapter) Ping() error {
	return mA.PingFn()
}
