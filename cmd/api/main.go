package main

import (
	"encoding/xml"
	"log"
	"os"

	"github.com/tiagocesar/yak-shop/handler"
	"github.com/tiagocesar/yak-shop/internal/models"
	"github.com/tiagocesar/yak-shop/internal/yak"
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

	yakSvc := yak.NewService(w.Yaks)

	h := handler.NewHttpServer(yakSvc)

	log.Printf("HTTP server starting on port 8080")
	h.ConfigureAndServe("8080")

	log.Println("HTTP server exiting")
	os.Exit(0)
}
