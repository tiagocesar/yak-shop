package handler

import (
	"github.com/tiagocesar/yak-shop/internal/models"
)

type stockHandlerResponse struct {
	Milk  float32 `json:"milk"`
	Skins int     `json:"skins"`
}

type herdHandlerResponse struct {
	Name string  `json:"name"`
	Age  float32 `json:"age"`
}

func toHerdHandlerResponse(herd []models.Yak) []herdHandlerResponse {
	result := make([]herdHandlerResponse, len(herd))

	for i := range herd {
		result[i] = herdHandlerResponse{
			Name: herd[i].Name,
			Age:  float32(herd[i].AgeInDays) / 100,
		}
	}

	return result
}
