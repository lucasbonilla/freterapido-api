package app

import (
	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	dbPostgresP ports.Postgres
	loggerP     ports.Logger
}

func NewAdapter(dbPostgresP ports.Postgres, loggerP ports.Logger) *Adapter {
	return &Adapter{
		dbPostgresP: dbPostgresP,
		loggerP:     loggerP,
	}
}

func (aA *Adapter) Run() {
	aA.dbPostgresP.InitConn()
	err := aA.dbPostgresP.Ping()
	if err != nil {
		aA.loggerP.Errorf("erro ao realizar ping na base de dados", err)
		return
	}
}
