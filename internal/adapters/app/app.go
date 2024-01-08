package app

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	dbPostgres ports.Postgres
	router     ports.Http
	config     ports.Config
	logger     ports.Logger
}

func NewAdapter(dbPostgres ports.Postgres, router ports.Http, configP ports.Config, logger ports.Logger) *Adapter {
	return &Adapter{
		dbPostgres: dbPostgres,
		router:     router,
		config:     configP,
		logger:     logger,
	}
}

func (aA *Adapter) Run() {
	aA.dbPostgres.InitConn()
	err := aA.dbPostgres.Ping()
	if err != nil {
		aA.logger.Errorf("erro ao realizar ping na base de dados", err)
		return
	}
	aA.router.InitializeRoutes()
	err = aA.router.Serve(fmt.Sprintf(":%s", aA.config.GetServerPort()))
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}