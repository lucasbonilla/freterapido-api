package remove_response

type Response struct {
	Dispatchers []Dispatchers `json:"dispatchers"`
}

type Carrier struct {
	Reference        int    `json:"reference"`
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
}

type DeliveryTime struct {
	Days          int    `json:"days,omitempty"`
	Hours         int    `json:"hours,omitempty"`
	Minutes       int    `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimated_date"`
}

type Weights struct {
	Real  float64 `json:"real"`
	Cubed float64 `json:"cubed,omitempty"`
	Used  float64 `json:"used"`
}

type SubTotal1 struct {
	Daily           int `json:"daily,omitempty"`
	Collect         int `json:"collect,omitempty"`
	Dispatch        int `json:"dispatch,omitempty"`
	Delivery        int `json:"delivery,omitempty"`
	Ferry           int `json:"ferry,omitempty"`
	Suframa         int `json:"suframa,omitempty"`
	Tas             int `json:"tas,omitempty"`
	SecCat          int `json:"sec_cat,omitempty"`
	Dat             int `json:"dat,omitempty"`
	AdValorem       int `json:"ad_valorem,omitempty"`
	Ademe           int `json:"ademe,omitempty"`
	Gris            int `json:"gris,omitempty"`
	Emex            int `json:"emex,omitempty"`
	Interior        int `json:"interior,omitempty"`
	Capatazia       int `json:"capatazia,omitempty"`
	River           int `json:"river,omitempty"`
	RiverInsurance  int `json:"river_insurance,omitempty"`
	Toll            int `json:"toll,omitempty"`
	Other           int `json:"other,omitempty"`
	OtherPerProduct int `json:"other_per_product,omitempty"`
}

type SubTotal2 struct {
	Trt        int `json:"trt,omitempty"`
	Tda        int `json:"tda,omitempty"`
	Tde        int `json:"tde,omitempty"`
	Scheduling int `json:"scheduling,omitempty"`
}

type SubTotal3 struct {
	Icms int `json:"icms,omitempty"`
}

type Composition struct {
	FreightWeight       float64   `json:"freight_weight"`
	FreightWeightExcess float64   `json:"freight_weight_excess"`
	FreightWeightVolume float64   `json:"freight_weight_volume"`
	FreightVolume       float64   `json:"freight_volume"`
	FreightMinimum      float64   `json:"freight_minimum"`
	FreightInvoice      float64   `json:"freight_invoice"`
	SubTotal1           SubTotal1 `json:"sub_total1"`
	SubTotal2           SubTotal2 `json:"sub_total2"`
	SubTotal3           SubTotal3 `json:"sub_total3"`
}

type OriginalDeliveryTime struct {
	Days          int    `json:"days,omitempty"`
	Hours         int    `json:"hours,omitempty"`
	Minutes       int    `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimated_date,omitempty"`
}

type Offers struct {
	Offer                int                  `json:"offer"`
	SimulationType       int                  `json:"simulation_type"`
	Carrier              Carrier              `json:"carrier"`
	Service              string               `json:"service"`
	ServiceCode          string               `json:"service_code,omitempty"`
	ServiceDescription   string               `json:"service_description,omitempty"`
	DeliveryTime         DeliveryTime         `json:"delivery_time"`
	Expiration           string               `json:"expiration"`
	CostPrice            float64              `json:"cost_price"`
	FinalPrice           float64              `json:"final_price"`
	Weights              Weights              `json:"weights"`
	Composition          Composition          `json:"composition,omitempty"`
	OriginalDeliveryTime OriginalDeliveryTime `json:"original_delivery_time,omitempty"`
	Identifier           string               `json:"identifier,omitempty"`
	DeliveryNote         string               `json:"delivery_note,omitempty"`
	HomeDelivery         bool                 `json:"home_delivery,omitempty"`
}

type Volumes struct {
	Category      string  `json:"category"`
	Sku           string  `json:"sku,omitempty"`
	Tag           string  `json:"tag,omitempty"`
	Description   string  `json:"description,omitempty"`
	Amount        int     `json:"amount"`
	Width         float64 `json:"width"`
	Height        float64 `json:"height"`
	Length        float64 `json:"length"`
	UnitaryWeight float64 `json:"unitary_weight"`
	UnitaryPrice  float64 `json:"unitary_price"`
	AmountVolumes float64 `json:"amount_volumes,omitempty"`
	Consolidate   bool    `json:"consolidate,omitempty"`
	Overlaid      bool    `json:"overlaid,omitempty"`
	Rotate        bool    `json:"rotate,omitempty"`
	Items         []any   `json:"items,omitempty"`
}

type Dispatchers struct {
	ID                         string    `json:"id"`
	RequestID                  string    `json:"request_id"`
	RegisteredNumberShipper    string    `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string    `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int       `json:"zipcode_origin"`
	Offers                     []Offers  `json:"offers,omitempty"`
	Volumes                    []Volumes `json:"volumes,omitempty"`
}
