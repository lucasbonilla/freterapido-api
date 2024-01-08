package ports

import (
	apiR "github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/api"
)

type Core interface {
	ValidateAPIRequest(APIRequest apiR.Request) []string
}
