package api

import (
	apiFRResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/frerapidoapi"
)

func NewResponse(apiResp *apiFRResp.Response) *Response {
	response := &Response{}
	for _, offer := range apiResp.Dispatchers[0].Offers {
		response.Carrier = append(response.Carrier, Carrier{
			Name:     offer.Carrier.Name,
			Service:  offer.Service,
			Deadline: offer.DeliveryTime.EstimatedDate,
			Price:    offer.CostPrice,
		})
	}

	return response
}

type Response struct {
	Carrier []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}
