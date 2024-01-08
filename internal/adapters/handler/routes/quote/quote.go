package quote

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
	APIReq "github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/api"
	APIFRReq "github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/frerapidoapi"
	APIResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/api"
	APIFRResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/frerapidoapi"
)

type Adapter struct {
	db      ports.Postgres
	message ports.Message
	core    ports.Core
	config  ports.Config
}

func NewAdapter(db ports.Postgres, message ports.Message, core ports.Core, config ports.Config) *Adapter {
	return &Adapter{
		db:      db,
		message: message,
		core:    core,
		config:  config,
	}
}

func (qA *Adapter) Quote(ctx *gin.Context) {
	request := APIReq.Request{}
	if err := ctx.BindJSON(&request); err != nil {
		qA.message.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := qA.core.ValidateAPIRequest(request); len(err) != 0 {
		qA.message.SendErrors(ctx, http.StatusBadRequest, err)
		return
	}

	reqFR := APIFRReq.NewRequest(request, qA.config.GetFreterapidoAPI())
	body, err := json.Marshal(reqFR)
	if err != nil {
		qA.message.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	payload := bytes.NewBuffer(body)
	resp, err := http.Post("https://sp.freterapido.com/api/v3/quote/simulate", "application/json", payload)
	if err != nil {
		qA.message.SendError(ctx, resp.StatusCode, err.Error())
	}

	APIresp := new(APIFRResp.Response)
	json.NewDecoder(resp.Body).Decode(APIresp)
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		qA.message.SendError(ctx, http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	respAPI := APIResp.NewResponse(APIresp)
	qA.message.SendSuccessWithCustomKey(ctx, "carrier", "quote", respAPI.Carrier)
}
