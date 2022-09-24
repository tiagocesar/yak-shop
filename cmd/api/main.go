package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/tiagocesar/yak-shop/internal/models"
)

const (
	minShavingAge = 101
	maxAge        = 1000
)

func main() {
	file, err := os.ReadFile("./herd-sample.xml")
	if err != nil {
		log.Fatal(err)
	}

	var w models.WrapperImport
	err = xml.Unmarshal(file, &w)
	if err != nil {
		log.Fatal(err)
	}

	herd := toHerd(w.Yaks)

	currentDay := 14 // Should be sanitized to not go above 1.000
	var totalMilk float32
	var totalWool int

	for i := range herd {
		yak := &herd[i]

		for day := 0; day < currentDay; day++ {
			if yak.Dead {
				break
			}

			// Milk
			producedMilk := 50 - (float32(yak.AgeInDays) * 0.03)

			if producedMilk > 0 {
				totalMilk += producedMilk
			}

			// Wool
			if yak.AgeInDays >= minShavingAge && yak.NextShave < day {
				totalWool++

				yak.NextShave = 8 + int(float32(yak.AgeInDays)*0.01)
			}

			// Ageing
			yak.AgeInDays += 1
			if yak.AgeInDays >= maxAge {
				yak.Dead = true
			}
		}
	}

	for _, yak := range herd {
		fmt.Println(yak)
	}

	fmt.Printf("%.3f liters of milk\n", totalMilk)
	fmt.Printf("%d skins of wool\n", totalWool)
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
