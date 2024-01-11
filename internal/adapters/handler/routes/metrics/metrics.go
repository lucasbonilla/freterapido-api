package metrics

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/api"
)

type Adapter struct {
	db      ports.Postgres
	http    ports.Http
	message ports.Message
	core    ports.Core
	config  ports.Config
	utils   ports.Utils
	logger  ports.Logger
}

func NewAdapter(db ports.Postgres, http ports.Http, message ports.Message, core ports.Core, config ports.Config,
	utils ports.Utils, logger ports.Logger) *Adapter {
	return &Adapter{
		db:      db,
		http:    http,
		message: message,
		core:    core,
		config:  config,
		utils:   utils,
		logger:  logger,
	}
}

func (mA *Adapter) Metrics(ctx *gin.Context) {
	var limit *int
	var offset *int
	var limitConv int
	lastQuotes := ctx.Query("last_quotes")
	page := ctx.Query("page")
	if lastQuotes != "" {
		limitConv, err := strconv.Atoi(lastQuotes)
		if err != nil {

		}
		limit = &limitConv
	}
	if page != "" && lastQuotes != "" {
		page, err := strconv.Atoi(page)
		if err != nil {

		}
		if page >= 1 {
			calcOffset := (limitConv * page) - limitConv
			offset = &calcOffset
		}
	}
	var metrics api.Metrics

	metrics.QuotesQuantity, _ = mA.db.GetNumberOfQuotes(limit, offset)
	metrics.TotalQuotesPrice, _ = mA.db.GetTotalQuotes(limit, offset)
	metrics.TotalQuotesAveragePrice, _ = mA.db.GetAverageQuotes(limit, offset)
	metrics.TotalQuotesCheapestPrice, _ = mA.db.GetCheapestQuotes(limit, offset)
	metrics.TotalQuoteMostExpensivePrice, _ = mA.db.GetMostExpensiveQuotes(limit, offset)
	mA.message.SendSuccessWithCustomKey(ctx, "metrics", "metrics", metrics)
}
