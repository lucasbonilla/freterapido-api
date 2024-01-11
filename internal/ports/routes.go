package ports

import "github.com/gin-gonic/gin"

type Quote interface {
	Quote(ctx *gin.Context)
}
