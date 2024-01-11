package freterapido

type FreterapidoAPI struct {
	BaseURL          string
	APIVersion       string
	RegisteredNumber string
	Token            string
	PlatformCode     string
	Zipcode          int
}

func NewConfig(baseURL string, apiVersion string, registeredNumber string, token string, platformCode string,
	zipcode int) *FreterapidoAPI {
	return &FreterapidoAPI{
		BaseURL:          baseURL,
		APIVersion:       apiVersion,
		RegisteredNumber: registeredNumber,
		Token:            token,
		PlatformCode:     platformCode,
		Zipcode:          zipcode,
	}
}
