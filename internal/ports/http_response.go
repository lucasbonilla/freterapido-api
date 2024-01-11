package ports

import "net/http"

type Res interface {
	SetResponse(res *http.Response)
	Close() error
}
