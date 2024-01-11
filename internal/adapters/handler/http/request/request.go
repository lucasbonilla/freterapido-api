package request

import "net/http"

type Adapter struct {
	req *http.Request
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (rA *Adapter) NewRequestWithContext() {

}
