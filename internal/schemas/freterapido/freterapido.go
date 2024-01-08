package freterapido

type FreterapidoAPI struct {
	BaseURL          string
	APIVersion       string
	RegisteredNumber string
	Token            string
	PlatformCode     string
	Zipcode          int
}

func NewConfig(baseURL string, APIVersion string, registeredNumber string, token string, platformCode string, zipcode int) *FreterapidoAPI {
	return &FreterapidoAPI{
		BaseURL:          baseURL,
		APIVersion:       APIVersion,
		RegisteredNumber: registeredNumber,
		Token:            token,
		PlatformCode:     platformCode,
		Zipcode:          zipcode,
	}
}
