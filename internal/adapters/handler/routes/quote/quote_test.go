package quote_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/config"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/db/postgres"
	httpH "github.com/lucasbonilla/freterapido-api/internal/adapters/handler/http"
	message "github.com/lucasbonilla/freterapido-api/internal/adapters/handler/message/response"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/handler/routes/quote"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/logger"
	"github.com/lucasbonilla/freterapido-api/internal/adapters/utils"
	"github.com/lucasbonilla/freterapido-api/internal/core"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/api"
	"github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

func TestQuote(t *testing.T) {
	var quoteP ports.Quote
	var mockedPostgres ports.Postgres
	var mockedHttp ports.Http
	var messageP ports.Message
	var coreP ports.Core
	var configP ports.Config
	var mockedUtils ports.Utils
	var logger ports.Logger
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
		AddCarrierFn: func(offers []freterapidoapi.Offers) error {
			return nil
		},
		AddQuoteFn: func(offers []freterapidoapi.Offers) error {
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
		ReadAllFn: func(body io.ReadCloser) ([]byte, error) {
			bytes := []byte(`{"dispatchers":[{"id":"659f32170544e6a757e47876","requestId":"659f32170544e6a757e47875","createdAt":"2024-01-11T00:11:03.147688484Z","registeredNumberShipper":"25438296000158","registeredNumberDispatcher":"25438296000158","zipcodeOrigin":29161376,"offers":[{"offer":1,"carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/34028316000103.png","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"PAC","serviceCode":"03298","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-01-16T00:00:00Z","costPrice":60.27,"finalPrice":60.27,"weights":{"real":13},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"identifier":"03298","homeDelivery":true},{"offer":2,"carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/34028316000103.png","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"SEDEX","serviceCode":"03220","deliveryTime":{"days":20,"estimatedDate":"2024-02-08"},"expiration":"2024-01-16T00:00:00Z","costPrice":151.89,"finalPrice":151.89,"weights":{"real":13},"originalDeliveryTime":{"days":10,"estimatedDate":"2024-01-25"},"identifier":"03220","homeDelivery":true},{"offer":3,"tableReference":"648a1ea756ea03bb5404ab4e","carrier":{"name":"RAPIDÃO FR (TESTE)","registeredNumber":"32964513000109","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/32964513000109.jpg","reference":"355","companyName":"TRANSPORTADORA RAPIDÃO FR (TESTE)"},"service":"teste2","deliveryTime":{"days":15,"hours":2,"estimatedDate":"2024-02-01"},"expiration":"2024-02-10T00:11:03.165115253Z","costPrice":158.38,"finalPrice":158.38,"weights":{"real":13,"cubed":24,"used":24},"originalDeliveryTime":{"days":5,"hours":2,"estimatedDate":"2024-01-18"},"homeDelivery":true},{"offer":4,"tableReference":"64df3e41007d99375df11e99","carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/correios/correios-pac.jpg","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"PAC","serviceCode":"03298","serviceDescription":"PAC CONTRATO AG","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-02-10T00:11:03.165124923Z","costPrice":170.63,"finalPrice":170.63,"weights":{"real":13,"used":17},"correios":{"declaredValue":true},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"homeDelivery":true},{"offer":5,"tableReference":"6462325329f1c8607fb8e54c","carrier":{"name":"JADLOG","registeredNumber":"04884082001107","stateInscription":"90421928-29","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/04884082001107.png","reference":"701","companyName":"JADLOG LOGISTICA S.A"},"service":".PACKAGE","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-02-10T00:11:03.165118471Z","costPrice":181.96,"finalPrice":181.96,"weights":{"real":13,"cubed":13.36,"used":13.36},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"homeDelivery":true},{"offer":6,"tableReference":"646b59b451f2b9d5942d250a","carrier":{"name":"BTU BRASPRESS","registeredNumber":"48740351002702","stateInscription":"103898530","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002702.png","reference":"474","companyName":"BRASPRESS TRANSPORTES URGENTES LTDA"},"service":"Normal","deliveryTime":{"days":27,"hours":7,"minutes":41,"estimatedDate":"2024-02-19"},"expiration":"2024-02-10T00:11:03.165120095Z","costPrice":290.04,"finalPrice":290.04,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":17,"hours":7,"minutes":41,"estimatedDate":"2024-02-05"},"homeDelivery":true},{"offer":7,"tableReference":"64df3e3e007d99375df11e98","carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/correios/correios-sedex.jpg","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"SEDEX","serviceCode":"03220","serviceDescription":"SEDEX CONTRATO AG","deliveryTime":{"days":20,"estimatedDate":"2024-02-08"},"expiration":"2024-02-10T00:11:03.165122762Z","costPrice":453.68,"finalPrice":453.68,"weights":{"real":13,"used":17},"correios":{"declaredValue":true},"originalDeliveryTime":{"days":10,"estimatedDate":"2024-01-25"},"homeDelivery":true},{"offer":8,"tableReference":"653f9b848e49d599a98531fd","carrier":{"name":"PRESSA FR (TESTE)","registeredNumber":"48740351002370","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002370.png","reference":"346","companyName":"PRESSA FR TRANSPORTES URGENTES (TESTE)"},"service":"Normal","deliveryTime":{"days":11,"estimatedDate":"2024-01-26"},"expiration":"2024-02-10T00:11:03.165107064Z","costPrice":1599.39,"finalPrice":1599.39,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":1,"estimatedDate":"2024-01-12"},"homeDelivery":true},{"offer":9,"tableReference":"657320732da2c72259d57935","carrier":{"name":"PRESSA FR (TESTE)","registeredNumber":"48740351002370","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002370.png","reference":"346","companyName":"PRESSA FR TRANSPORTES URGENTES (TESTE)"},"service":"Normal","deliveryTime":{"days":11,"estimatedDate":"2024-01-26"},"expiration":"2024-02-10T00:11:03.165113572Z","costPrice":1599.39,"finalPrice":1599.39,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":1,"estimatedDate":"2024-01-12"},"homeDelivery":true}]}]}`)
			return bytes, nil
		},
		JSONUnmarshalFn: func(data []byte, v any) error {
			return json.Unmarshal(data, v)
		},
	}

	quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
		mockedUtils, logger)

	var ctx *gin.Context

	quoteP.Quote(ctx)
}

func TestQuoteError(t *testing.T) {
	t.Run("error on bind json", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{}
		mockedHttp = &httpH.MockedAdapter{}
		coreP = &core.MockedAdapter{}

		mockedUtils = &utils.MockedAdapter{
			BindJSONFn: func(ctx *gin.Context, obj any) (interface{}, error) {
				return nil, errors.New("ocorreu um erro")
			},
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
	t.Run("error on data assertion", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{}
		mockedHttp = &httpH.MockedAdapter{}
		coreP = &core.MockedAdapter{}

		mockedUtils = &utils.MockedAdapter{
			BindJSONFn: func(ctx *gin.Context, obj any) (interface{}, error) {
				var interfaceReturn interface{}
				return interfaceReturn, nil
			},
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
	t.Run("error on ValidateAPIRequest", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{}
		mockedHttp = &httpH.MockedAdapter{}
		coreP = &core.MockedAdapter{
			ValidateAPIRequestFn: func(apiRequest api.Request) []string {
				var validateErrors []string
				validateErrors = append(validateErrors, "Campo inválido")
				return validateErrors
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
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
	t.Run("error on JSONMarshal", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{}
		mockedHttp = &httpH.MockedAdapter{}
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
				var byteArray []byte
				return byteArray, errors.New("Ocorreu um erro")
			},
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
	t.Run("error on Do", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{}
		mockedHttp = &httpH.MockedAdapter{
			DoFn: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("Ocorreu um erro")
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
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
	t.Run("error on ReadAll", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{}
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
			ReadAllFn: func(body io.ReadCloser) ([]byte, error) {
				var byteArray []byte
				return byteArray, errors.New("Ocorreu um erro")
			},
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
	t.Run("error on JSONUnmarshal", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{}
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
			ReadAllFn: func(body io.ReadCloser) ([]byte, error) {
				bytes := []byte(`{"dispatchers":[{"id":"659f32170544e6a757e47876","requestId":"659f32170544e6a757e47875","createdAt":"2024-01-11T00:11:03.147688484Z","registeredNumberShipper":"25438296000158","registeredNumberDispatcher":"25438296000158","zipcodeOrigin":29161376,"offers":[{"offer":1,"carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/34028316000103.png","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"PAC","serviceCode":"03298","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-01-16T00:00:00Z","costPrice":60.27,"finalPrice":60.27,"weights":{"real":13},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"identifier":"03298","homeDelivery":true},{"offer":2,"carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/34028316000103.png","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"SEDEX","serviceCode":"03220","deliveryTime":{"days":20,"estimatedDate":"2024-02-08"},"expiration":"2024-01-16T00:00:00Z","costPrice":151.89,"finalPrice":151.89,"weights":{"real":13},"originalDeliveryTime":{"days":10,"estimatedDate":"2024-01-25"},"identifier":"03220","homeDelivery":true},{"offer":3,"tableReference":"648a1ea756ea03bb5404ab4e","carrier":{"name":"RAPIDÃO FR (TESTE)","registeredNumber":"32964513000109","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/32964513000109.jpg","reference":"355","companyName":"TRANSPORTADORA RAPIDÃO FR (TESTE)"},"service":"teste2","deliveryTime":{"days":15,"hours":2,"estimatedDate":"2024-02-01"},"expiration":"2024-02-10T00:11:03.165115253Z","costPrice":158.38,"finalPrice":158.38,"weights":{"real":13,"cubed":24,"used":24},"originalDeliveryTime":{"days":5,"hours":2,"estimatedDate":"2024-01-18"},"homeDelivery":true},{"offer":4,"tableReference":"64df3e41007d99375df11e99","carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/correios/correios-pac.jpg","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"PAC","serviceCode":"03298","serviceDescription":"PAC CONTRATO AG","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-02-10T00:11:03.165124923Z","costPrice":170.63,"finalPrice":170.63,"weights":{"real":13,"used":17},"correios":{"declaredValue":true},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"homeDelivery":true},{"offer":5,"tableReference":"6462325329f1c8607fb8e54c","carrier":{"name":"JADLOG","registeredNumber":"04884082001107","stateInscription":"90421928-29","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/04884082001107.png","reference":"701","companyName":"JADLOG LOGISTICA S.A"},"service":".PACKAGE","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-02-10T00:11:03.165118471Z","costPrice":181.96,"finalPrice":181.96,"weights":{"real":13,"cubed":13.36,"used":13.36},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"homeDelivery":true},{"offer":6,"tableReference":"646b59b451f2b9d5942d250a","carrier":{"name":"BTU BRASPRESS","registeredNumber":"48740351002702","stateInscription":"103898530","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002702.png","reference":"474","companyName":"BRASPRESS TRANSPORTES URGENTES LTDA"},"service":"Normal","deliveryTime":{"days":27,"hours":7,"minutes":41,"estimatedDate":"2024-02-19"},"expiration":"2024-02-10T00:11:03.165120095Z","costPrice":290.04,"finalPrice":290.04,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":17,"hours":7,"minutes":41,"estimatedDate":"2024-02-05"},"homeDelivery":true},{"offer":7,"tableReference":"64df3e3e007d99375df11e98","carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/correios/correios-sedex.jpg","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"SEDEX","serviceCode":"03220","serviceDescription":"SEDEX CONTRATO AG","deliveryTime":{"days":20,"estimatedDate":"2024-02-08"},"expiration":"2024-02-10T00:11:03.165122762Z","costPrice":453.68,"finalPrice":453.68,"weights":{"real":13,"used":17},"correios":{"declaredValue":true},"originalDeliveryTime":{"days":10,"estimatedDate":"2024-01-25"},"homeDelivery":true},{"offer":8,"tableReference":"653f9b848e49d599a98531fd","carrier":{"name":"PRESSA FR (TESTE)","registeredNumber":"48740351002370","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002370.png","reference":"346","companyName":"PRESSA FR TRANSPORTES URGENTES (TESTE)"},"service":"Normal","deliveryTime":{"days":11,"estimatedDate":"2024-01-26"},"expiration":"2024-02-10T00:11:03.165107064Z","costPrice":1599.39,"finalPrice":1599.39,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":1,"estimatedDate":"2024-01-12"},"homeDelivery":true},{"offer":9,"tableReference":"657320732da2c72259d57935","carrier":{"name":"PRESSA FR (TESTE)","registeredNumber":"48740351002370","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002370.png","reference":"346","companyName":"PRESSA FR TRANSPORTES URGENTES (TESTE)"},"service":"Normal","deliveryTime":{"days":11,"estimatedDate":"2024-01-26"},"expiration":"2024-02-10T00:11:03.165113572Z","costPrice":1599.39,"finalPrice":1599.39,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":1,"estimatedDate":"2024-01-12"},"homeDelivery":true}]}]}`)
				return bytes, nil
			},
			JSONUnmarshalFn: func(data []byte, v any) error {
				return errors.New("Ocorreu um erro")
			},
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
	t.Run("error on db store", func(t *testing.T) {
		var quoteP ports.Quote
		var mockedPostgres ports.Postgres
		var mockedHttp ports.Http
		var messageP ports.Message
		var coreP ports.Core
		var configP ports.Config
		var mockedUtils ports.Utils
		var loggerP ports.Logger
		messageP = &message.MockedAdapter{
			SendSuccessFn:              func(ctx *gin.Context, op string, data interface{}) { return },
			SendSuccessWithCustomKeyFn: func(ctx *gin.Context, key, op string, data interface{}) { return },
			SendErrorFn:                func(ctx *gin.Context, code int, msg string) { return },
			SendErrorsFn:               func(ctx *gin.Context, code int, msg []string) { return },
		}
		configP = config.NewMockedAdapter()
		mockedPostgres = &postgres.MockedAdapter{
			AddCarrierFn: func(offers []freterapidoapi.Offers) error {
				return errors.New("Ocorreu um erro")
			},
			AddQuoteFn: func(offers []freterapidoapi.Offers) error {
				return errors.New("Ocorreu um erro")
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
			ReadAllFn: func(body io.ReadCloser) ([]byte, error) {
				bytes := []byte(`{"dispatchers":[{"id":"659f32170544e6a757e47876","requestId":"659f32170544e6a757e47875","createdAt":"2024-01-11T00:11:03.147688484Z","registeredNumberShipper":"25438296000158","registeredNumberDispatcher":"25438296000158","zipcodeOrigin":29161376,"offers":[{"offer":1,"carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/34028316000103.png","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"PAC","serviceCode":"03298","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-01-16T00:00:00Z","costPrice":60.27,"finalPrice":60.27,"weights":{"real":13},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"identifier":"03298","homeDelivery":true},{"offer":2,"carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/34028316000103.png","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"SEDEX","serviceCode":"03220","deliveryTime":{"days":20,"estimatedDate":"2024-02-08"},"expiration":"2024-01-16T00:00:00Z","costPrice":151.89,"finalPrice":151.89,"weights":{"real":13},"originalDeliveryTime":{"days":10,"estimatedDate":"2024-01-25"},"identifier":"03220","homeDelivery":true},{"offer":3,"tableReference":"648a1ea756ea03bb5404ab4e","carrier":{"name":"RAPIDÃO FR (TESTE)","registeredNumber":"32964513000109","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/32964513000109.jpg","reference":"355","companyName":"TRANSPORTADORA RAPIDÃO FR (TESTE)"},"service":"teste2","deliveryTime":{"days":15,"hours":2,"estimatedDate":"2024-02-01"},"expiration":"2024-02-10T00:11:03.165115253Z","costPrice":158.38,"finalPrice":158.38,"weights":{"real":13,"cubed":24,"used":24},"originalDeliveryTime":{"days":5,"hours":2,"estimatedDate":"2024-01-18"},"homeDelivery":true},{"offer":4,"tableReference":"64df3e41007d99375df11e99","carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/correios/correios-pac.jpg","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"PAC","serviceCode":"03298","serviceDescription":"PAC CONTRATO AG","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-02-10T00:11:03.165124923Z","costPrice":170.63,"finalPrice":170.63,"weights":{"real":13,"used":17},"correios":{"declaredValue":true},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"homeDelivery":true},{"offer":5,"tableReference":"6462325329f1c8607fb8e54c","carrier":{"name":"JADLOG","registeredNumber":"04884082001107","stateInscription":"90421928-29","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/04884082001107.png","reference":"701","companyName":"JADLOG LOGISTICA S.A"},"service":".PACKAGE","deliveryTime":{"days":24,"estimatedDate":"2024-02-14"},"expiration":"2024-02-10T00:11:03.165118471Z","costPrice":181.96,"finalPrice":181.96,"weights":{"real":13,"cubed":13.36,"used":13.36},"originalDeliveryTime":{"days":14,"estimatedDate":"2024-01-31"},"homeDelivery":true},{"offer":6,"tableReference":"646b59b451f2b9d5942d250a","carrier":{"name":"BTU BRASPRESS","registeredNumber":"48740351002702","stateInscription":"103898530","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002702.png","reference":"474","companyName":"BRASPRESS TRANSPORTES URGENTES LTDA"},"service":"Normal","deliveryTime":{"days":27,"hours":7,"minutes":41,"estimatedDate":"2024-02-19"},"expiration":"2024-02-10T00:11:03.165120095Z","costPrice":290.04,"finalPrice":290.04,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":17,"hours":7,"minutes":41,"estimatedDate":"2024-02-05"},"homeDelivery":true},{"offer":7,"tableReference":"64df3e3e007d99375df11e98","carrier":{"name":"CORREIOS","registeredNumber":"34028316000103","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/correios/correios-sedex.jpg","reference":"281","companyName":"EMPRESA BRASILEIRA DE CORREIOS E TELEGRAFOS"},"service":"SEDEX","serviceCode":"03220","serviceDescription":"SEDEX CONTRATO AG","deliveryTime":{"days":20,"estimatedDate":"2024-02-08"},"expiration":"2024-02-10T00:11:03.165122762Z","costPrice":453.68,"finalPrice":453.68,"weights":{"real":13,"used":17},"correios":{"declaredValue":true},"originalDeliveryTime":{"days":10,"estimatedDate":"2024-01-25"},"homeDelivery":true},{"offer":8,"tableReference":"653f9b848e49d599a98531fd","carrier":{"name":"PRESSA FR (TESTE)","registeredNumber":"48740351002370","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002370.png","reference":"346","companyName":"PRESSA FR TRANSPORTES URGENTES (TESTE)"},"service":"Normal","deliveryTime":{"days":11,"estimatedDate":"2024-01-26"},"expiration":"2024-02-10T00:11:03.165107064Z","costPrice":1599.39,"finalPrice":1599.39,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":1,"estimatedDate":"2024-01-12"},"homeDelivery":true},{"offer":9,"tableReference":"657320732da2c72259d57935","carrier":{"name":"PRESSA FR (TESTE)","registeredNumber":"48740351002370","stateInscription":"ISENTO","logo":"https://s3.amazonaws.com/public.prod.freterapido.uploads/transportadora/foto-perfil/48740351002370.png","reference":"346","companyName":"PRESSA FR TRANSPORTES URGENTES (TESTE)"},"service":"Normal","deliveryTime":{"days":11,"estimatedDate":"2024-01-26"},"expiration":"2024-02-10T00:11:03.165113572Z","costPrice":1599.39,"finalPrice":1599.39,"weights":{"real":13,"cubed":16,"used":16},"originalDeliveryTime":{"days":1,"estimatedDate":"2024-01-12"},"homeDelivery":true}]}]}`)
				return bytes, nil
			},
			JSONUnmarshalFn: func(data []byte, v any) error {
				return json.Unmarshal(data, v)
			},
		}
		loggerP = logger.NewAdapter(configP)

		quoteP = quote.NewAdapter(mockedPostgres, mockedHttp, messageP, coreP, configP,
			mockedUtils, loggerP)

		var ctx *gin.Context

		quoteP.Quote(ctx)
	})
}
