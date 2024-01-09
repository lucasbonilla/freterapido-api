package postgres

type MockedAdapter struct {
	InitConnFn func() error
	PingFn     func() error
}

func (mA *MockedAdapter) InitConn() error {
	return mA.InitConnFn()
}

func (mA *MockedAdapter) Ping() error {
	return mA.PingFn()
}
