package ports

import "github.com/gin-gonic/gin"

type Message interface {
	SendSuccess(ctx *gin.Context, op string, data interface{})
	SendSuccessWithCustomKey(ctx *gin.Context, key string, op string, data interface{})
	SendError(ctx *gin.Context, code int, msg string)
	SendErrors(ctx *gin.Context, code int, msg []string)
}
