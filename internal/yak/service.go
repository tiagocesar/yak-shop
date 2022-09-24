package yak

import (
	"github.com/tiagocesar/yak-shop/internal/models"
)

type Service struct {
	herd []models.Yak // In a real scenario it would be a data store
}

func NewService(yakImport []models.YakImport) *Service {
	herd := toHerd(yakImport)

	return &Service{herd: herd}
}

func (s *Service) Process(day int) (float32, int) {
	if day > models.MaxYakAge {
		// adjusting to the maximum possible age for a yak to be alive
		day = models.MaxYakAge
	}

	var totalMilk float32
	var totalWool int

	for i := range s.herd {
		yak := &s.herd[i]

		for i := 0; i < day; i++ {
			if yak.Dead {
				break
			}

			totalMilk += yak.Milk()
			totalWool += yak.Shave(i)

			yak.Age()
		}
	}

	return totalMilk, totalWool
}

func toHerd(imports []models.YakImport) []models.Yak {
	result := make([]models.Yak, len(imports))

	for i := 0; i < len(imports); i++ {
		result[i] = models.Yak{
			Name:      imports[i].Name,
			AgeInDays: int(imports[i].Age * 100),
			Sex:       imports[i].Sex,
			Dead:      false,
		}
	}

	return result
}
