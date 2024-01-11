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

func (rA *Adapter) ValidateAPIRequest(apiRequest apiR.Request) []string {
	errs := make([]string, 0)
	_, err := strconv.Atoi(apiRequest.Recipient.Address.Zipcode)
	if err != nil || len(apiRequest.Recipient.Address.Zipcode) != 8 {
		errs = append(errs, "ZipCode inválido")
	}
	if len(apiRequest.Volumes) < 1 {
		errs = append(errs, "Pelo menos um volume deve ser enviado")
	}
	for index, volume := range apiRequest.Volumes {
		switch {
		case volume.Category == nil:
			errs = append(errs, fmt.Sprintf("Category deve ser informado para o produto %v", index+1))
		case volume.Amount == nil:
			errs = append(errs, fmt.Sprintf("Amount deve ser informado para o produto %v", index+1))
		case volume.UnitaryWeight == nil:
			errs = append(errs, fmt.Sprintf("UnitaryWeight deve ser informado para o produto %v", index+1))
		case volume.Price == nil:
			errs = append(errs, fmt.Sprintf("Price deve ser informado para o produto %v", index+1))
		case volume.Height == nil:
			errs = append(errs, fmt.Sprintf("Height deve ser informado para o produto %v", index+1))
		case volume.Width == nil:
			errs = append(errs, fmt.Sprintf("Width deve ser informado para o produto %v", index+1))
		case volume.Length == nil:
			errs = append(errs, fmt.Sprintf("Length deve ser informado para o produto %v", index+1))
		}
	}

	return errs
}
