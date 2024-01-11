package postgres

import (
	APIResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/api"
	freterapidoapi "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

type MockedAdapter struct {
	InitConnFn               func() error
	CloseFn                  func() error
	PingFn                   func() error
	AddCarrierFn             func(offers []freterapidoapi.Offers) error
	AddQuoteFn               func(offers []freterapidoapi.Offers) error
	GetNumberOfQuotesFn      func(limit *int, offset *int) (*APIResp.QuotesQuantity, error)
	GetTotalQuotesFn         func(limit *int, offset *int) (APIResp.TotalQuotesPrice, error)
	GetAverageQuotesfn       func(limit *int, offset *int) (APIResp.TotalQuotesAveragePrice, error)
	GetCheapestQuotesFn      func(limit *int, offset *int) (APIResp.TotalQuotesCheapestPrice, error)
	GetMostExpensiveQuotesFn func(limit *int, offset *int) (APIResp.TotalQuoteMostExpensivePrice, error)
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

func (mA *MockedAdapter) GetNumberOfQuotes(limit *int, offset *int) (APIResp.QuotesQuantity, error) {
	return mA.GetNumberOfQuotes(limit, offset)
}

func (mA *MockedAdapter) GetTotalQuotes(limit *int, offset *int) (APIResp.TotalQuotesPrice, error) {
	return mA.GetTotalQuotesFn(limit, offset)
}

func (mA *MockedAdapter) GetAverageQuotes(limit *int, offset *int) (APIResp.TotalQuotesAveragePrice, error) {
	return mA.GetAverageQuotes(limit, offset)
}

func (mA *MockedAdapter) GetCheapestQuotes(limit *int, offset *int) (APIResp.TotalQuotesCheapestPrice, error) {
	return mA.GetCheapestQuotesFn(limit, offset)
}

func (mA *MockedAdapter) GetMostExpensiveQuotes(limit *int, offset *int) (APIResp.TotalQuoteMostExpensivePrice, error) {
	return mA.GetMostExpensiveQuotesFn(limit, offset)
}
