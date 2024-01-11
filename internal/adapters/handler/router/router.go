package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	gin   *gin.Engine
	quote ports.Quote
}

func NewAdapter(quote ports.Quote) *Adapter {
	return &Adapter{
		quote: quote,
	}
}

func (gA *Adapter) InitializeRoutes() {
	gA.gin = gin.Default()

	basePath := "/api/v1"
	v1 := gA.gin.Group(basePath)
	{
		v1.POST("quote", gA.quote.Quote)
	}
}

func (gA *Adapter) Serve(listenAdr string) error {
	return gA.gin.Run(listenAdr)
}
