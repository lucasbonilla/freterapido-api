package utils

import (
	"io"

	"github.com/gin-gonic/gin"
)

type MockedAdapter struct {
	BindJSONFn    func(ctx *gin.Context, obj any) (interface{}, error)
	JSONMarshalFn func(v any) ([]byte, error)
	JSONDecodeFn  func(r io.Reader, v any) (interface{}, error)
}

func (mA *MockedAdapter) BindJSON(ctx *gin.Context, obj any) (interface{}, error) {
	return mA.BindJSONFn(ctx, obj)
}
func (mA *MockedAdapter) JSONMarshal(v any) ([]byte, error) {
	return mA.JSONMarshalFn(v)
}
func (mA *MockedAdapter) JSONDecode(r io.Reader, v any) (interface{}, error) {
	return mA.JSONDecodeFn(r, v)
}
