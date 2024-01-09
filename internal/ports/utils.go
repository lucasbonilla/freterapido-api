package ports

import (
	"io"

	"github.com/gin-gonic/gin"
)

type Utils interface {
	BindJSON(ctx *gin.Context, obj any) (interface{}, error)
	JSONMarshal(v any) ([]byte, error)
	JSONDecode(r io.Reader, v any) (interface{}, error)
}
