package utils

import (
	"io"

	"github.com/gin-gonic/gin"
)

type MockedAdapter struct {
	BindJSONFn      func(ctx *gin.Context, obj any) (interface{}, error)
	JSONMarshalFn   func(v any) ([]byte, error)
	ReadAllFn       func(body io.ReadCloser) ([]byte, error)
	JSONUnmarshalFn func(data []byte, v any) error
}

func (mA *MockedAdapter) BindJSON(ctx *gin.Context, obj any) (interface{}, error) {
	return mA.BindJSONFn(ctx, obj)
}
func (mA *MockedAdapter) JSONMarshal(v any) ([]byte, error) {
	return mA.JSONMarshalFn(v)
}

func (mA *MockedAdapter) ReadAll(body io.ReadCloser) ([]byte, error) {
	return mA.ReadAllFn(body)
}
func (mA *MockedAdapter) JSONUnmarshal(data []byte, v any) error {
	return mA.JSONUnmarshalFn(data, v)
}
