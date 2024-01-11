package client

import (
	"net/http"
)

type Adapter struct {
	cli *http.Client
}

func NewAdapter() *Adapter {
	return &Adapter{
		cli: &http.Client{},
	}
}

func (cA *Adapter) Do(req *http.Request) (*http.Response, error) {
	return cA.cli.Do(req)
}
