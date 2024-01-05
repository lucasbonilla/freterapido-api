package ports

import "github.com/lucasbonilla/freterapido-api/internal/schemas/db"

type Config interface {
	GetDB() *db.Config
	GetServerPort() string
	RunType() string
}
