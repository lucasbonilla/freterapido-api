package app

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	dbPostgres ports.Postgres
	router     ports.Router
	config     ports.Config
	logger     ports.Logger
}

func NewAdapter(dbPostgres ports.Postgres, router ports.Router, configP ports.Config, logger ports.Logger) *Adapter {
	return &Adapter{
		dbPostgres: dbPostgres,
		router:     router,
		config:     configP,
		logger:     logger,
	}
}

func (aA *Adapter) Run() {
	// inicializa configurações da base de dados
	aA.dbPostgres.InitConn()
	err := aA.dbPostgres.Ping()
	if err != nil {
		aA.logger.Errorf("erro ao realizar ping na base de dados", err)

		return
	}

	// inicializa rotas
	aA.router.InitializeRoutes()
	err = aA.router.Serve(fmt.Sprintf(":%s", aA.config.GetServerPort()))
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
	aA.dbPostgres.Close()
}
