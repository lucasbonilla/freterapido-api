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
	APIFRResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

type Adapter struct {
	db      ports.Postgres
	http    ports.Http
	message ports.Message
	core    ports.Core
	config  ports.Config
	utils   ports.Utils
	logger  ports.Logger
}

func NewAdapter(db ports.Postgres, http ports.Http, message ports.Message, core ports.Core, config ports.Config,
	utils ports.Utils, logger ports.Logger) *Adapter {
	return &Adapter{
		db:      db,
		http:    http,
		message: message,
		core:    core,
		config:  config,
		utils:   utils,
		logger:  logger,
	}
}

func (qA *Adapter) Quote(ctx *gin.Context) {
	var request *APIReq.Request
	requestToBind := APIReq.Request{}
	reqBind, err := qA.utils.BindJSON(ctx, &requestToBind)
	if err != nil {
		qA.logger.Errorf("erro json malformatado: %v", err.Error())
		qA.message.SendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	// realiza o type assertion da requisição para poder ser testada e garantida a existência dos dados.
	var okAssertion bool
	if request, okAssertion = reqBind.(*APIReq.Request); !okAssertion {
		qA.logger.Error("erro json malformatado")
		qA.message.SendError(ctx, http.StatusInternalServerError,
			"Ocorreu um erro ao realizar bind da request")

		return
	}

	// valida a existência de todos os dados esperados na requisição
	if err := qA.core.ValidateAPIRequest(*request); len(err) != 0 {
		qA.logger.Errorf("erro de validação: %v", err)
		qA.message.SendErrors(ctx, http.StatusBadRequest, err)

		return
	}

	reqFR := APIFRReq.NewRequest(*request, qA.config.GetFreterapidoAPI())
	body, err := qA.utils.JSONMarshal(reqFR)
	if err != nil {
		qA.logger.Errorf("erro json marshal: %v", err.Error())
		qA.message.SendError(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	payload := bytes.NewReader(body)
	req, _ := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf(
			"https://%s/%s/quote/simulate",
			qA.config.GetFreterapidoAPI().BaseURL,
			qA.config.GetFreterapidoAPI().APIVersion),
		payload)

	resp, err := qA.http.Do(req)
	if err != nil {
		qA.logger.Errorf("erro ao realizar requisição: %v", err.Error())
		qA.message.SendError(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	qA.http.SetResponse(resp)
	defer qA.http.Close()

	APIresp := APIFRResp.Response{}
	bytes, err := qA.utils.ReadAll(resp.Body)
	if err != nil {
		qA.logger.Errorf("erro ao ler o corpo da requisição: %v", err.Error())
		qA.message.SendError(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	err = qA.utils.JSONUnmarshal(bytes, &APIresp)
	if err != nil {
		qA.logger.Errorf("erro ao realizar unmarshal da requisição: %v", err.Error())
		qA.message.SendError(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	respAPI := APIResp.NewResponse(&APIresp)

	// realiza a persistência dos dados na base de dados
	qA.persistCarrierQuote(err, APIresp)
	qA.message.SendSuccessWithCustomKey(ctx, "carrier", "quote", respAPI.Carrier)
}

func (qA *Adapter) persistCarrierQuote(err error, apiResp APIFRResp.Response) {
	err = qA.db.AddCarrier(apiResp.Dispatchers[0].Offers)
	if err != nil {
		qA.logger.Errorf("ocorreu um erro ao armazenar transportadoras na base de dados", err.Error())
	}

	err = qA.db.AddQuote(apiResp.Dispatchers[0].Offers)
	if err != nil {
		qA.logger.Errorf("ocorreu um erro ao armazenar cotações na base de dados", err.Error())
	}
}
