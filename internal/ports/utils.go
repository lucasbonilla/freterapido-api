package ports

import (
	"io"

	"github.com/gin-gonic/gin"
)

type Utils interface {
	BindJSON(ctx *gin.Context, obj any) (interface{}, error)
	JSONMarshal(v any) ([]byte, error)
	ReadAll(body io.ReadCloser) ([]byte, error)
	JSONUnmarshal(data []byte, v any) error
}
