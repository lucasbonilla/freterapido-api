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
	GetDBFn             func() *db.Config
	GetServerPortFn     func() string
	RunTypeFn           func() string
	GetFreterapidoAPIFn func() *freterapido.FreterapidoAPI
}

func NewMockedAdapter() *MockedAdapter {
	return &MockedAdapter{
		APIConfig:      api.NewConfig("8080", "test"),
		DBConfig:       db.NewConfig("local-test", "15432", "db-test", "db-test", "postgres-test"),
		APIFreterapido: freterapido.NewConfig("base-url-test", "v-test", "abcd1234", "dcba4321", "1234abcd", 10101000),
	}
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
