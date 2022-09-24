package yak

import (
	"github.com/tiagocesar/yak-shop/internal/models"
)

type Service struct {
	herd []models.Yak // In a real scenario it would be a data store
}

func NewService(herd []models.Yak) *Service {
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
