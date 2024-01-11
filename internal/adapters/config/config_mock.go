package config

import (
	"github.com/lucasbonilla/freterapido-api/internal/schemas/api"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/db"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/freterapido"
)

type MockedAdapter struct {
	APIConfig           *api.Config
	DBConfig            *db.Config
	APIFreterapido      *freterapido.FreterapidoAPI
	InitConfigFn        func() error
	GetDBFn             func() *db.Config
	GetServerPortFn     func() string
	RunTypeFn           func() string
	GetFreterapidoAPIFn func() *freterapido.FreterapidoAPI
}

func (mA *MockedAdapter) InitConfig() error {
	return mA.InitConfigFn()
}

func (mA *MockedAdapter) GetDB() *db.Config {
	return mA.DBConfig
}

func (mA *MockedAdapter) GetServerPort() string {
	return mA.APIConfig.APIPort
}

func (mA *MockedAdapter) RunType() string {
	return mA.APIConfig.RunType
}

func (mA *MockedAdapter) GetFreterapidoAPI() *freterapido.FreterapidoAPI {
	return mA.APIFreterapido
}
