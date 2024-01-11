package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	gin     *gin.Engine
	quote   ports.Quote
	metrics ports.Metrics
}

func NewAdapter(quote ports.Quote, metrics ports.Metrics) *Adapter {
	return &Adapter{
		quote:   quote,
		metrics: metrics,
	}
}

func (gA *Adapter) InitializeRoutes() {
	gA.gin = gin.Default()

	basePath := "/api/v1"
	v1 := gA.gin.Group(basePath)
	{
		v1.POST("quote", gA.quote.Quote)
		v1.GET("metrics", gA.metrics.Metrics)
	}
}

func (gA *Adapter) Serve(listenAdr string) error {
	return gA.gin.Run(listenAdr)
}
