package api

type Request struct {
	Recipient Recipient `json:"recipient,omitempty"`
	Volumes   []Volumes `json:"volumes,omitempty"`
}
type Address struct {
	Zipcode string `json:"zipcode,omitempty"`
}
type Recipient struct {
	Address Address `json:"address,omitempty"`
}
type Volumes struct {
	Category      *int     `json:"category,omitempty"`
	Amount        *int     `json:"amount,omitempty"`
	UnitaryWeight *int     `json:"unitary_weight,omitempty"`
	Price         *int     `json:"price,omitempty"`
	Sku           string   `json:"sku,omitempty"`
	Height        *float64 `json:"height,omitempty"`
	Width         *float64 `json:"width,omitempty"`
	Length        *float64 `json:"length,omitempty"`
}
