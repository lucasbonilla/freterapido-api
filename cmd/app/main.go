package main

import (
	"github.com/lucasbonilla/freterapido-api/internal/adapters/app"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/config"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/db/postgres"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/logger"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

func main() {
	var configP ports.Config
	var loggerP ports.Logger
	var dbPostgresP ports.Postgres

	var appP ports.App

	configP = config.NewAdpter()
	loggerP = logger.NewAdapter(configP)
	dbPostgresP = postgres.NewAdapter(configP)

	appP = app.NewAdapter(dbPostgresP, loggerP)
	appP.Run()
}
