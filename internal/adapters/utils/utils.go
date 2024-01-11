package utils

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (uA *Adapter) BindJSON(ctx *gin.Context, obj any) (interface{}, error) {
	err := ctx.BindJSON(obj)

	return obj, err
}

func (uA *Adapter) JSONMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (uA *Adapter) ReadAll(body io.ReadCloser) ([]byte, error) {
	return io.ReadAll(body)
}

func (uA *Adapter) JSONUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
