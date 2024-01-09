package quote

import (
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/config"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/db/postgres"
	httpH "github.com/lucasbonilla/freterapido-api/internal/adapters/handler/http"
	message "github.com/lucasbonilla/freterapido-api/internal/adapters/handler/message/response"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/utils"
	"github.com/lucasbonilla/freterapido-api/internal/core"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/api"
	APIFRResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/frerapidoapi"
)

func TestQuote(t *testing.T) {
	var quoteP ports.Quote
	var mockedPostgres ports.Postgres
	var mockedHttp ports.Http
	var messageP ports.Message
	var coreP ports.Core
	var configP ports.Config
	var mockedUtils ports.Utils
	messageP = &message.MockedAdapter{
		SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
		SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
		SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
		SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
	}
	configP = config.NewMockedAdapter()
	mockedPostgres = &postgres.MockedAdapter{
		InitConnFn: func() error {
			return nil
		},
		PingFn: func() error {
			return nil
		},
	}
	mockedHttp = &httpH.MockedAdapter{
		DoFn: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, nil
		},
		CloseFn: func() error {
			return nil
		},
	}
	coreP = &core.MockedAdapter{
		ValidateAPIRequestFn: func(apiRequest api.Request) []string {
			var err []string
			return err
		},
	}

	var category, amount, unitaryWeight, price int
	var height, width, length float64

	category = 7
	amount = 1
	unitaryWeight = 5
	price = 349

	height = 0.2
	width = 0.2
	length = 0.2

	mockedUtils = &utils.MockedAdapter{
		BindJSONFn: func(ctx *gin.Context, obj any) (interface{}, error) {
			apiRequest := api.Request{
				Recipient: api.Recipient{
					Address: api.Address{
						Zipcode: "97070220"},
				},
				Volumes: []api.Volumes{
					{
						Category:      &category,
						Amount:        &amount,
						UnitaryWeight: &unitaryWeight,
						Price:         &price,
						Sku:           "teste",
						Height:        &height,
						Width:         &width,
						Length:        &length}},
			}
			var response interface{} = &apiRequest
			return response, nil
		},
		JSONMarshalFn: func(v any) ([]byte, error) {
			return []byte{}, nil
		},
		JSONDecodeFn: func(r io.Reader, v any) (interface{}, error) {
			FRResponse := APIFRResp.Response{
				Dispatchers: []APIFRResp.Dispatchers{
					{
						ID:                         "1234567890",
						RequestID:                  "",
						RegisteredNumberShipper:    "",
						RegisteredNumberDispatcher: "",
						ZipcodeOrigin:              0,
						Offers: []APIFRResp.Offers{
							{
								Offer:          1,
								SimulationType: 0,
								Carrier: APIFRResp.Carrier{
									Reference:        "111",
									Name:             "carrier",
									RegisteredNumber: "",
									StateInscription: "",
									Logo:             "https://logo.com",
								},
								Service:            "Correios",
								ServiceCode:        "",
								ServiceDescription: "",
								DeliveryTime: APIFRResp.DeliveryTime{
									Days:          0,
									Hours:         0,
									Minutes:       0,
									EstimatedDate: "",
								},
								Expiration: "2024-02-08T17:25:11.288174565Z",
								CostPrice:  0,
								FinalPrice: 0,
								Weights: APIFRResp.Weights{
									Real:  13,
									Cubed: 13.36,
									Used:  13.36,
								},
								Composition: APIFRResp.Composition{
									FreightWeight:       0,
									FreightWeightExcess: 0,
									FreightWeightVolume: 0,
									FreightVolume:       0,
									FreightMinimum:      0,
									FreightInvoice:      0,
									SubTotal1:           APIFRResp.SubTotal1{},
									SubTotal2:           APIFRResp.SubTotal2{},
									SubTotal3:           APIFRResp.SubTotal3{},
								},
								OriginalDeliveryTime: APIFRResp.OriginalDeliveryTime{},
								Identifier:           "",
								DeliveryNote:         "",
								HomeDelivery:         false,
							},
							{
								Offer:          1,
								SimulationType: 0,
								Carrier: APIFRResp.Carrier{
									Reference:        "222",
									Name:             "carrier",
									RegisteredNumber: "",
									StateInscription: "",
									Logo:             "https://logo.com",
								},
								Service:            "transportadora_a",
								ServiceCode:        "",
								ServiceDescription: "",
								DeliveryTime: APIFRResp.DeliveryTime{
									Days:          0,
									Hours:         0,
									Minutes:       0,
									EstimatedDate: "",
								},
								Expiration: "2024-02-08T17:25:11.288174565Z",
								CostPrice:  0,
								FinalPrice: 0,
								Weights: APIFRResp.Weights{
									Real:  13,
									Cubed: 13.36,
									Used:  13.36,
								},
								Composition: APIFRResp.Composition{
									FreightWeight:       0,
									FreightWeightExcess: 0,
									FreightWeightVolume: 0,
									FreightVolume:       0,
									FreightMinimum:      0,
									FreightInvoice:      0,
									SubTotal1:           APIFRResp.SubTotal1{},
									SubTotal2:           APIFRResp.SubTotal2{},
									SubTotal3:           APIFRResp.SubTotal3{},
								},
								OriginalDeliveryTime: APIFRResp.OriginalDeliveryTime{},
								Identifier:           "",
								DeliveryNote:         "",
								HomeDelivery:         false,
							},
						},
						Volumes: []APIFRResp.Volumes{}}}}
			var response interface{} = &FRResponse
			return response, nil
		},
	}

	quoteP = NewAdapter(mockedPostgres,
		mockedHttp, messageP, coreP, configP, mockedUtils)

	var ctx *gin.Context

	quoteP.Quote(ctx)
	// NewAdapter(dbPostgresP, httpP, messageP, coreP, configP, utilsP)
}
