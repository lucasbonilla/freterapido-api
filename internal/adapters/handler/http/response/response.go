package response

import "net/http"

type Adapter struct {
	res *http.Response
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (r *Adapter) Close() error {
	return r.res.Body.Close()
}
