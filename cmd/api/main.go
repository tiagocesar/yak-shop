package main

import (
	"encoding/xml"
	"log"
	"os"

	"github.com/tiagocesar/yak-shop/internal/models"
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
}
