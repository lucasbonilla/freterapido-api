package metrics

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/api"
)

type Adapter struct {
	db      ports.Postgres
	message ports.Message
	core    ports.Core
	config  ports.Config
	utils   ports.Utils
	logger  ports.Logger
}

func NewAdapter(db ports.Postgres, message ports.Message, config ports.Config, utils ports.Utils,
	logger ports.Logger) *Adapter {
	return &Adapter{
		db:      db,
		message: message,
		config:  config,
		utils:   utils,
		logger:  logger,
	}
}

func (mA *Adapter) Metrics(ctx *gin.Context) {
	var limit *int
	var offset *int
	var limitConv int
	var err error
	lastQuotes := ctx.Query("last_quotes")
	page := ctx.Query("page")
	if lastQuotes != "" {
		limitConv, err = strconv.Atoi(lastQuotes)
		if err != nil {
			mA.logger.Errorf("erro ao capturar last_quotes: %v", err.Error())
			mA.message.SendError(ctx, http.StatusInternalServerError, err.Error())

			return
		}
		limit = &limitConv
	}
	if page != "" && lastQuotes != "" {
		page, err := strconv.Atoi(page)
		if err != nil {
			mA.logger.Errorf("erro ao capturar page: %v", err.Error())
			mA.message.SendError(ctx, http.StatusInternalServerError, err.Error())

			return
		}
		if page >= 1 {
			calcOffset := (limitConv * page) - limitConv
			offset = &calcOffset
		}
	}
	var metrics api.Metrics

	// cada método abaixo faz uma consulta diferente para cada tipo de resposta esperado
	// caso algum SQL tenha que ser alterado basta alterar na sua unidade ao invés de alterar um sql maior
	// que levaria mais tempo e mais esforço
	metrics.QuotesQuantity, _ = mA.db.GetNumberOfQuotes(limit, offset)
	metrics.TotalQuotesPrice, _ = mA.db.GetTotalQuotes(limit, offset)
	metrics.TotalQuotesAveragePrice, _ = mA.db.GetAverageQuotes(limit, offset)
	metrics.TotalQuotesCheapestPrice, _ = mA.db.GetCheapestQuotes(limit, offset)
	metrics.TotalQuoteMostExpensivePrice, _ = mA.db.GetMostExpensiveQuotes(limit, offset)
	mA.message.SendSuccessWithCustomKey(ctx, "metrics", "metrics", metrics)
}
