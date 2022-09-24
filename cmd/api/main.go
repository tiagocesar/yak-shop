package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/tiagocesar/yak-shop/internal/models"
)

func main() {
	file, err := os.ReadFile("./herd-sample.xml")
	if err != nil {
		log.Fatal(err)
	}

	var herd []models.Herd
	err = xml.Unmarshal(file, &herd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(herd)
}
