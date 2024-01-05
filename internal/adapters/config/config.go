package config

import (
	"github.com/lucasbonilla/freterapido-api/internal/schemas/api"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/db"
	"github.com/spf13/viper"
)

type Adapter struct {
	APIConfig *api.Config
	DBConfig  *db.Config
}

func NewAdpter() *Adapter {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")

	viper.SetConfigFile("../../config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil
		}
	}

	apiConfig := api.NewConfig(viper.GetString("api.port"), viper.GetString("run.type"))
	dbConfig := db.NewConfig(
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.pass"),
		viper.GetString("database.name"))

	return &Adapter{
		APIConfig: apiConfig,
		DBConfig:  dbConfig,
	}
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
