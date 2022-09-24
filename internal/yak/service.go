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

// Process does the ageing of herd while
// processing the production of milk and wool
func (s *Service) Process(day int) (float32, int) {
	if day > models.MaxYakAge {
		// adjusting to the maximum possible age for a yak to be alive
		day = models.MaxYakAge
	}

	var totalMilk float32
	var totalWool int

	// Making a copy so the original data is preserved
	herd := make([]models.Yak, len(s.herd))
	copy(herd, s.herd)

	for i := range herd {
		yak := &herd[i]

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

// GetHerdInfo does a "dry run" over the herd, ageing it
// and getting the information of the state after the specified
// days period has passed.
func (s *Service) GetHerdInfo(day int) []models.Yak {
	if day > models.MaxYakAge {
		// adjusting to the maximum possible age for a yak to be alive
		day = models.MaxYakAge
	}

	// Making a copy so the original data is preserved
	herd := make([]models.Yak, len(s.herd))
	copy(herd, s.herd)

	for i := range herd {
		yak := &herd[i]

		for i := 0; i < day; i++ {
			if yak.Dead {
				break
			}

			// We simulate shaving the yak because we want to know when it was last shaved
			yak.Shave(i)

			yak.Age()
		}
	}

	return herd
}

// toHerd converts from the format used to import the input file
// to a format we can work with internally
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
