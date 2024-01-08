package core

import (
	"fmt"
	"strconv"

	apiR "github.com/lucasbonilla/freterapido-api/internal/schemas/message/request/api"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (rA *Adapter) ValidateAPIRequest(APIRequest apiR.Request) []string {
	errs := make([]string, 0)
	_, err := strconv.Atoi(APIRequest.Recipient.Address.Zipcode)
	if err != nil || len(APIRequest.Recipient.Address.Zipcode) != 8 {
		errs = append(errs, "ZipCode inválido")
	}
	if len(APIRequest.Volumes) < 1 {
		errs = append(errs, "Pelo menos um volume deve ser enviado")
	}
	for i, volume := range APIRequest.Volumes {
		switch {
		case volume.Category == nil:
			errs = append(errs, fmt.Sprintf("Category deve ser informado para o produto %v", i+1))
		case volume.Amount == nil:
			errs = append(errs, fmt.Sprintf("Amount deve ser informado para o produto %v", i+1))
		case volume.UnitaryWeight == nil:
			errs = append(errs, fmt.Sprintf("UnitaryWeight deve ser informado para o produto %v", i+1))
		case volume.Price == nil:
			errs = append(errs, fmt.Sprintf("Price deve ser informado para o produto %v", i+1))
		case volume.Height == nil:
			errs = append(errs, fmt.Sprintf("Height deve ser informado para o produto %v", i+1))
		case volume.Width == nil:
			errs = append(errs, fmt.Sprintf("Width deve ser informado para o produto %v", i+1))
		case volume.Length == nil:
			errs = append(errs, fmt.Sprintf("Length deve ser informado para o produto %v", i+1))
		}
	}
	return errs
}
