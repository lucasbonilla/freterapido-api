package message

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (mA *Adapter) SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from route: %s successfull", op),
		"data":    data,
	})
}

func (mA *Adapter) SendSuccessWithCustomKey(ctx *gin.Context, key string, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from route: %s successfull", op),
		key:       data,
	})
}

func (mA *Adapter) SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func (mA *Adapter) SendErrors(ctx *gin.Context, code int, msg []string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
