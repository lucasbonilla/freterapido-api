package config

import (
	"github.com/lucasbonilla/freterapido-api/internal/schemas/api"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/db"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/freterapido"
	"github.com/spf13/viper"
)

type Adapter struct {
	APIConfig      *api.Config
	DBConfig       *db.Config
	APIFreterapido *freterapido.FreterapidoAPI
}

func NewAdpter() *Adapter {
	return &Adapter{}
}

func (cA *Adapter) InitConfig() error {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")

	viper.SetConfigFile("../../config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	apiConfig := api.NewConfig(
		viper.GetString("api.port"),
		viper.GetString("run.type"))

	dbConfig := db.NewConfig(
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.pass"),
		viper.GetString("database.name"))

	freterapidoAPI := freterapido.NewConfig(
		viper.GetString("freterapido_api.base_url"),
		viper.GetString("freterapido_api.api_version"),
		viper.GetString("freterapido_api.registered_number"),
		viper.GetString("freterapido_api.token"),
		viper.GetString("freterapido_api.platform_code"),
		viper.GetInt("freterapido_api.zipcode"),
	)

	cA.APIConfig = apiConfig
	cA.DBConfig = dbConfig
	cA.APIFreterapido = freterapidoAPI

	return nil
}

func (cA *Adapter) GetDB() *db.Config {
	return cA.DBConfig
}

func (cA *Adapter) GetServerPort() string {
	return cA.APIConfig.APIPort
}

func (cA *Adapter) RunType() string {
	return cA.APIConfig.RunType
}

func (cA *Adapter) GetFreterapidoAPI() *freterapido.FreterapidoAPI {
	return cA.APIFreterapido
}
