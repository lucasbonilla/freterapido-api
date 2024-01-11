package http

import (
	"net/http"

	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	cli ports.Cli
	res ports.Res
}

func NewAdapter(cli ports.Cli, res ports.Res) *Adapter {
	return &Adapter{
		cli: cli,
		res: res,
	}
}

func (httpA *Adapter) Do(req *http.Request) (*http.Response, error) {
	return httpA.cli.Do(req)
}

func (httpA *Adapter) SetResponse(res *http.Response) {
	httpA.res.SetResponse(res)
}

func (httpA *Adapter) Close() error {
	return httpA.res.Close()
}
