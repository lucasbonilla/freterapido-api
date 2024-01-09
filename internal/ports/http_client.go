package ports

import "net/http"

type Cli interface {
	Do(req *http.Request) (*http.Response, error)
}
