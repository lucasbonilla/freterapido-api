package postgres

import freterapidoapi "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"

type MockedAdapter struct {
	InitConnFn   func() error
	CloseFn      func() error
	PingFn       func() error
	AddCarrierFn func(offers []freterapidoapi.Offers) error
	AddQuoteFn   func(offers []freterapidoapi.Offers) error
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

func (mA *MockedAdapter) AddCarrier(offers []freterapidoapi.Offers) error {
	return mA.AddCarrierFn(offers)
}
func (mA *MockedAdapter) AddQuote(offers []freterapidoapi.Offers) error {
	return mA.AddQuoteFn(offers)
}
