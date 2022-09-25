package handler

import (
	"errors"
	"fmt"

	"github.com/tiagocesar/yak-shop/internal/models"
)

var (
	ErrCustomerNotSpecified = errors.New("please specify a customer")
	ErrNoGoodsSpecified     = errors.New("please specify at least a valid item for the order")
)

type stockHandlerResponse struct {
	Milk  float32 `json:"milk"`
	Skins int     `json:"skins"`
}

type herdHandlerResponse struct {
	Name          string `json:"name"`
	Age           string `json:"age"`
	AgeLastShaved string `json:"age-last-shaved"`
}

func toHerdHandlerResponse(herd []models.Yak) []herdHandlerResponse {
	result := make([]herdHandlerResponse, len(herd))

	for i := range herd {
		age := float32(herd[i].AgeInDays) / 100
		lastShaved := float32(herd[i].AgeLastShaved) / 100

		result[i] = herdHandlerResponse{
			Name:          herd[i].Name,
			Age:           fmt.Sprintf("%v", age),
			AgeLastShaved: fmt.Sprintf("%.1f", lastShaved),
		}
	}

	return result
}

type orderHandlerRequest struct {
	Customer string `json:"customer"`
	Order    struct {
		Milk  float32 `json:"milk"`
		Skins int     `json:"skins"`
	} `json:"order"`
}

func (o orderHandlerRequest) validate() error {
	if o.Customer == "" {
		return ErrCustomerNotSpecified
	}

	if o.Order.Milk == 0 && o.Order.Skins == 0 {
		return ErrNoGoodsSpecified
	}

	return nil
}

type orderHandlerResponse struct {
	Milk  float32 `json:"milk,omitempty"`
	Skins int     `json:"skins,omitempty"`
}
