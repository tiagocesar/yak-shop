package handler

import (
	"fmt"

	"github.com/tiagocesar/yak-shop/internal/models"
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
