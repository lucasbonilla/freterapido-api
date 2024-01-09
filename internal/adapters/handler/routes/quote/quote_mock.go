package quote

import "github.com/gin-gonic/gin"

type MockedAdapter struct {
	QuoteFn func(ctx *gin.Context)
}

func (mA *MockedAdapter) Quote(ctx *gin.Context) {
	return
}
