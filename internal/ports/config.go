package ports

import (
	"github.com/lucasbonilla/freterapido-api/internal/schemas/db"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/freterapido"
)

type Config interface {
	GetDB() *db.Config
	GetServerPort() string
	RunType() string
	GetFreterapidoAPI() *freterapido.FreterapidoAPI
}
