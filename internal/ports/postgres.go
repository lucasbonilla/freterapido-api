package ports

import (
	"github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

type Postgres interface {
	InitConn() error
	Close() error
	Ping() error
	AddCarrier(offers []freterapidoapi.Offers) error
	AddQuote(offers []freterapidoapi.Offers) error
}
