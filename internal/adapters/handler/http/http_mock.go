package http

import "net/http"

type MockedAdapter struct {
	DoFn          func(req *http.Request) (*http.Response, error)
	SetResponseFn func(res *http.Response)
	CloseFn       func() error
}

func (mA *MockedAdapter) Do(req *http.Request) (*http.Response, error) {
	return mA.DoFn(req)
}

func (mA *MockedAdapter) SetResponse(res *http.Response) {
	return
}

func (mA *MockedAdapter) Close() error {
	return mA.CloseFn()
}
