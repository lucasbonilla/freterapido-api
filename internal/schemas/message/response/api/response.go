package api

import (
	apiFRResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

func NewResponse(apiResp *apiFRResp.Response) *Response {
	response := &Response{}
	for _, offer := range apiResp.Dispatchers[0].Offers {
		response.Carrier = append(response.Carrier, Carrier{
			Name:     offer.Carrier.Name,
			Service:  offer.Service,
			Deadline: offer.DeliveryTime.Days,
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
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}

type Metrics struct {
	QuotesQuantity               QuotesQuantity               `json:"quotes_quantity"`
	TotalQuotesPrice             TotalQuotesPrice             `json:"total_quotes_price"`
	TotalQuotesAveragePrice      TotalQuotesAveragePrice      `json:"total_quotes_average_price"`
	TotalQuotesCheapestPrice     TotalQuotesCheapestPrice     `json:"total_quotes_cheapest_price"`
	TotalQuoteMostExpensivePrice TotalQuoteMostExpensivePrice `json:"total_quote_most_expensive_price"`
}

type QuotesQuantity struct {
	QuotesByQuantity []QuotesByQuantity `json:"quotes_by_quantity"`
}
type QuotesByQuantity struct {
	CarrierName    string `json:"carrier_name"`
	IDCarrier      int    `json:"id_carrier"`
	NumberOfQuotes int    `json:"number_of_quotes"`
}

type TotalQuotesPrice struct {
	TotalQuotesByPrice []TotalQuotesByPrice `json:"total_quotes_by_price"`
}
type TotalQuotesByPrice struct {
	CarrierName     string  `json:"carrier_name"`
	IDCarrier       int     `json:"id_carrier"`
	TotalPriceQuote float64 `json:"total_price_quote"`
}

type TotalQuotesAveragePrice struct {
	TotalQuotesByAveragePrice []TotalQuotesByAveragePrice `json:"total_quotes_by_average_price"`
}
type TotalQuotesByAveragePrice struct {
	CarrierName       string  `json:"carrier_name"`
	IDCarrier         int     `json:"id_carrier"`
	AveragePriceQuote float64 `json:"average_price_quote"`
}

type TotalQuotesCheapestPrice struct {
	TotalQuotesForCheapestPrice []TotalQuotesForCheapestPrice `json:"total_quotes_for_cheapest_price"`
}
type TotalQuotesForCheapestPrice struct {
	CarrierName        string  `json:"carrier_name"`
	IDCarrier          int     `json:"id_carrier"`
	PriceQuoteCheapest float64 `json:"price_quote_cheapest"`
}

type TotalQuoteMostExpensivePrice struct {
	TotalQuoteByMostExpensivePrice []TotalQuoteByMostExpensivePrice `json:"total_quote_by_most_expensive_price"`
}
type TotalQuoteByMostExpensivePrice struct {
	CarrierName             string  `json:"carrier_name"`
	IDCarrier               int     `json:"id_carrier"`
	PriceQuoteMostExpensive float64 `json:"price_quote_most_expensive"`
}
