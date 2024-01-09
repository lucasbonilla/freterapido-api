package ports

import "net/http"

type Http interface {
	Do(req *http.Request) (*http.Response, error)
	Close() error
}
