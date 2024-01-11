package ports

import (
	APIResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/api"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

type Postgres interface {
	InitConn() error
	Close() error
	Ping() error
	AddCarrier(offers []freterapidoapi.Offers) error
	AddQuote(offers []freterapidoapi.Offers) error
	GetNumberOfQuotes(limit *int, offset *int) (APIResp.QuotesQuantity, error)
	GetTotalQuotes(limit *int, offset *int) (APIResp.TotalQuotesPrice, error)
	GetAverageQuotes(limit *int, offset *int) (APIResp.TotalQuotesAveragePrice, error)
	GetCheapestQuotes(limit *int, offset *int) (APIResp.TotalQuotesCheapestPrice, error)
	GetMostExpensiveQuotes(limit *int, offset *int) (APIResp.TotalQuoteMostExpensivePrice, error)
}
