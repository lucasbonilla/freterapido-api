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

func (uA *Adapter) JSONDecode(r io.Reader, v any) (interface{}, error) {
	err := json.NewDecoder(r).Decode(v)
	return v, err
}
