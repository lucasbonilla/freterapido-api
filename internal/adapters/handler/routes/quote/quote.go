package quote

import (
	"bytes"
	"fmt"
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
	http    ports.Http
	message ports.Message
	core    ports.Core
	config  ports.Config
	utils   ports.Utils
}

func NewAdapter(db ports.Postgres, http ports.Http, message ports.Message, core ports.Core, config ports.Config, utils ports.Utils) *Adapter {
	return &Adapter{
		db:      db,
		http:    http,
		message: message,
		core:    core,
		config:  config,
		utils:   utils,
	}
}

func (qA *Adapter) Quote(ctx *gin.Context) {
	var request *APIReq.Request
	requestToBind := APIReq.Request{}
	var ok bool
	reqBind, err := qA.utils.BindJSON(ctx, &requestToBind)

	if request, ok = reqBind.(*APIReq.Request); !ok {
		qA.message.SendError(ctx, http.StatusInternalServerError,
			"Ocorreu um erro ao realizar bind da request")
	}
	fmt.Println(request)

	if err != nil {
		qA.message.SendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	if err := qA.core.ValidateAPIRequest(*request); len(err) != 0 {
		qA.message.SendErrors(ctx, http.StatusBadRequest, err)

		return
	}

	reqFR := APIFRReq.NewRequest(*request, qA.config.GetFreterapidoAPI())
	body, err := qA.utils.JSONMarshal(reqFR)
	if err != nil {
		qA.message.SendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	payload := bytes.NewReader(body)

	req, _ := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://sp.freterapido.com/api/v3/quote/simulate",
		payload)

	resp, err := qA.http.Do(req)
	if err != nil {
		qA.message.SendError(ctx, resp.StatusCode, err.Error())
	}

	APIresp := new(APIFRResp.Response)
	responseToDecode := APIFRResp.Response{}
	respDecode, err := qA.utils.JSONDecode(resp.Body, &responseToDecode)
	if err != nil {
		qA.message.SendError(ctx, http.StatusInternalServerError, err.Error())
	}

	defer qA.http.Close()

	if APIresp, ok = respDecode.(*APIFRResp.Response); !ok {
		qA.message.SendError(ctx, http.StatusInternalServerError,
			"Ocorreu um erro ao realizar bind da request")
	}

	respAPI := APIResp.NewResponse(APIresp)
	qA.message.SendSuccessWithCustomKey(ctx, "carrier", "quote", respAPI.Carrier)
}
