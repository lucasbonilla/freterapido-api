package freterapidoapi

type Response struct {
	Dispatchers []Dispatchers `json:"dispatchers"`
}

type Carrier struct {
	Reference        string `json:"reference"`
	Name             string `json:"name"`
	RegisteredNumber string `json:"registeredNumber"`
	StateInscription string `json:"stateInscription"`
	Logo             string `json:"logo"`
}

type DeliveryTime struct {
	Days          int    `json:"days,omitempty"`
	Hours         int    `json:"hours,omitempty"`
	Minutes       int    `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimatedDate"`
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
	SecCat          int `json:"SecCat,omitempty"`
	Dat             int `json:"dat,omitempty"`
	AdValorem       int `json:"AdValorem,omitempty"`
	Ademe           int `json:"ademe,omitempty"`
	Gris            int `json:"gris,omitempty"`
	Emex            int `json:"emex,omitempty"`
	Interior        int `json:"interior,omitempty"`
	Capatazia       int `json:"capatazia,omitempty"`
	River           int `json:"river,omitempty"`
	RiverInsurance  int `json:"RiverInsurance,omitempty"`
	Toll            int `json:"toll,omitempty"`
	Other           int `json:"other,omitempty"`
	OtherPerProduct int `json:"OtherPerProduct,omitempty"`
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
	FreightWeight       float64   `json:"FreightWeight"`
	FreightWeightExcess float64   `json:"FreightWeightExcess"`
	FreightWeightVolume float64   `json:"FreightWeightVolume"`
	FreightVolume       float64   `json:"FreightVolume"`
	FreightMinimum      float64   `json:"FreightMinimum"`
	FreightInvoice      float64   `json:"FreightInvoice"`
	SubTotal1           SubTotal1 `json:"SubTotal1"`
	SubTotal2           SubTotal2 `json:"SubTotal2"`
	SubTotal3           SubTotal3 `json:"SubTotal3"`
}

type OriginalDeliveryTime struct {
	Days          int    `json:"days,omitempty"`
	Hours         int    `json:"hours,omitempty"`
	Minutes       int    `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimatedDate,omitempty"`
}

type Offers struct {
	Offer                int                  `json:"offer"`
	SimulationType       int                  `json:"simulationType"`
	Carrier              Carrier              `json:"carrier"`
	Service              string               `json:"service"`
	ServiceCode          string               `json:"serviceCode,omitempty"`
	ServiceDescription   string               `json:"serviceDescription,omitempty"`
	DeliveryTime         DeliveryTime         `json:"deliveryTime"`
	Expiration           string               `json:"expiration"`
	CostPrice            float64              `json:"costPrice"`
	FinalPrice           float64              `json:"finalPrice"`
	Weights              Weights              `json:"weights"`
	Composition          Composition          `json:"composition,omitempty"`
	OriginalDeliveryTime OriginalDeliveryTime `json:"originalDeliveryTime,omitempty"`
	Identifier           string               `json:"identifier,omitempty"`
	DeliveryNote         string               `json:"DeliveryNote,omitempty"`
	HomeDelivery         bool                 `json:"homeDelivery,omitempty"`
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
	UnitaryWeight float64 `json:"UnitaryWeight"`
	UnitaryPrice  float64 `json:"UnitaryPrice"`
	AmountVolumes float64 `json:"AmountVolumes,omitempty"`
	Consolidate   bool    `json:"consolidate,omitempty"`
	Overlaid      bool    `json:"overlaid,omitempty"`
	Rotate        bool    `json:"rotate,omitempty"`
	Items         []any   `json:"items,omitempty"`
}

type Dispatchers struct {
	ID                         string    `json:"id"`
	RequestID                  string    `json:"RequestID"`
	RegisteredNumberShipper    string    `json:"RegisteredNumberShipper"`
	RegisteredNumberDispatcher string    `json:"RegisteredNumberDispatcher"`
	ZipcodeOrigin              int       `json:"ZipcodeOrigin"`
	Offers                     []Offers  `json:"offers,omitempty"`
	Volumes                    []Volumes `json:"volumes,omitempty"`
}
