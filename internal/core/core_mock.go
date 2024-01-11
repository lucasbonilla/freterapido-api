package core

import apiR "github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/api"

type MockedAdapter struct {
	ValidateAPIRequestFn func(apiRequest apiR.Request) []string
}

func (mA *MockedAdapter) ValidateAPIRequest(apiRequest apiR.Request) []string {
	return mA.ValidateAPIRequestFn(apiRequest)
}
