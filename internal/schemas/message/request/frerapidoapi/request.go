package freterapidoapi

import (
	"strconv"

	"github.com/lucasbonilla/freterapido-api/internal/schemas/freterapido"
	apiR "github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/api"
)

func NewRequest(request apiR.Request, freterapidoAPI *freterapido.FreterapidoAPI) *Request {
	zipcode, _ := strconv.Atoi(request.Recipient.Address.Zipcode)
	req := &Request{
		Shipper: Shipper{
			RegisteredNumber: freterapidoAPI.RegisteredNumber,
			Token:            freterapidoAPI.Token,
			PlatformCode:     freterapidoAPI.PlatformCode,
		},
		Recipient: Recipient{
			Type:    0,
			Country: "BRA",
			Zipcode: zipcode,
		},
		Dispatchers: []Dispatchers{
			{
				RegisteredNumber: freterapidoAPI.RegisteredNumber,
				Zipcode:          freterapidoAPI.Zipcode,
				Volumes:          []Volumes{},
			},
		},
		SimulationType: []int{0},
	}

	for _, volume := range request.Volumes {
		req.Dispatchers[0].Volumes = append(
			req.Dispatchers[0].Volumes, Volumes{
				Category:      strconv.Itoa(*volume.Category),
				Amount:        *volume.Amount,
				UnitaryWeight: float64(*volume.UnitaryWeight),
				UnitaryPrice:  float64(*volume.Price),
				Sku:           volume.Sku,
				Height:        *volume.Height,
				Width:         *volume.Width,
				Length:        *volume.Length,
			})
	}

	return req
}

type Request struct {
	Shipper        Shipper       `json:"shipper"`
	Recipient      Recipient     `json:"recipient,omitempty"`
	Dispatchers    []Dispatchers `json:"dispatchers,omitempty"`
	Channel        string        `json:"channel,omitempty"`
	Filter         int           `json:"filter,omitempty"`
	Limit          int           `json:"limit,omitempty"`
	Identification string        `json:"identification,omitempty"`
	Reverse        bool          `json:"reverse,omitempty"`
	SimulationType []int         `json:"simulation_type"`
	Returns        Returns       `json:"returns,omitempty"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type             int    `json:"type"`
	RegisteredNumber string `json:"registered_number,omitempty"`
	StateInscription string `json:"state_inscription,omitempty"`
	Country          string `json:"country"`
	Zipcode          int    `json:"zipcode"`
}

type Volumes struct {
	Amount        int     `json:"amount"`
	AmountVolumes int     `json:"amount_volumes"`
	Category      string  `json:"category,omitempty"`
	Sku           string  `json:"sku,omitempty"`
	Tag           string  `json:"tag,omitempty"`
	Description   string  `json:"description,omitempty"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Consolidate   bool    `json:"consolidate,omitempty"`
	Overlaid      bool    `json:"overlaid,omitempty"`
	Rotate        bool    `json:"rotate,omitempty"`
}

type Dispatchers struct {
	RegisteredNumber string    `json:"registered_number"`
	Zipcode          int       `json:"zipcode"`
	TotalPrice       float64   `json:"total_price,omitempty"`
	Volumes          []Volumes `json:"volumes"`
}

type Returns struct {
	Composition  bool `json:"composition,omitempty"`
	Volumes      bool `json:"volumes,omitempty"`
	AppliedRules bool `json:"applied_rules,omitempty"`
}
