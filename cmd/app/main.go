package main

import (
	"github.com/lucasbonilla/freterapido-api/internal/adapters/app"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/config"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/db/postgres"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/http"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/http/client"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/http/response"
	messageResponse "github.com/lucasbonilla/freterapido-api/internal/adapters/handler/message/response"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/router"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/routes/metrics"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/routes/quote"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/logger"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/utils"
	"github.com/lucasbonilla/freterapido-api/internal/core"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

func main() {
	var configP ports.Config
	var loggerP ports.Logger
	var dbPostgresP ports.Postgres

	var utilsP ports.Utils

	var httpP ports.Http
	var httpCli ports.Cli
	var httpRes ports.Res

	var coreP ports.Core
	var messageP ports.Message
	var quoteP ports.Quote
	var metricsP ports.Metrics
	var routerP ports.Router

	var appP ports.App

	// inicialização de todos os adapters
	coreP = core.NewAdapter()
	configP = config.NewAdpter()
	loggerP = logger.NewAdapter(configP)
	dbPostgresP = postgres.NewAdapter(configP)

	utilsP = utils.NewAdapter()

	httpCli = client.NewAdapter()
	httpRes = response.NewAdapter()
	httpP = http.NewAdapter(httpCli, httpRes)

	messageP = messageResponse.NewAdapter()
	quoteP = quote.NewAdapter(dbPostgresP, httpP, messageP, coreP, configP, utilsP,
		loggerP)
	metricsP = metrics.NewAdapter(dbPostgresP, messageP, configP, utilsP, loggerP)
	routerP = router.NewAdapter(quoteP, metricsP)

	appP = app.NewAdapter(dbPostgresP, routerP, configP, loggerP)
	appP.Run()
}
