package main

import (
	"github.com/lucasbonilla/freterapido-api/internal/adapters/app"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/config"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/db/postgres"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/message/message"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/router"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/routes/quote"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/logger"
	"github.com/lucasbonilla/freterapido-api/internal/core"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

func main() {
	var configP ports.Config
	var loggerP ports.Logger
	var dbPostgresP ports.Postgres

	var coreP ports.Core
	var messageP ports.Message
	var quoteP ports.Quote
	var routerP ports.Http

	var appP ports.App

	coreP = core.NewAdapter()
	configP = config.NewAdpter()
	loggerP = logger.NewAdapter(configP)
	dbPostgresP = postgres.NewAdapter(configP)

	messageP = message.NewAdapter()
	quoteP = quote.NewAdapter(dbPostgresP, messageP, coreP, configP)
	routerP = router.NewAdapter(quoteP)

	appP = app.NewAdapter(dbPostgresP, routerP, configP, loggerP)
	appP.Run()
}
