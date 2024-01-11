package message

import "github.com/gin-gonic/gin"

type MockedAdapter struct {
	SendSuccessFn              func(ctx *gin.Context, op string, data interface{})
	SendSuccessWithCustomKeyFn func(ctx *gin.Context, key string, op string, data interface{})
	SendErrorFn                func(ctx *gin.Context, code int, msg string)
	SendErrorsFn               func(ctx *gin.Context, code int, msg []string)
}

func (mA *MockedAdapter) SendSuccess(ctx *gin.Context, op string, data interface{}) {
	return
}

func (mA *MockedAdapter) SendSuccessWithCustomKey(ctx *gin.Context, key string, op string, data interface{}) {
	return
}

func (mA *MockedAdapter) SendError(ctx *gin.Context, code int, msg string) {
	return
}

func (mA *MockedAdapter) SendErrors(ctx *gin.Context, code int, msg []string) {
	return
}
