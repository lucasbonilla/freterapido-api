package ports

import "net/http"

type Http interface {
	Do(req *http.Request) (*http.Response, error)
	SetResponse(res *http.Response)
	Close() error
}
