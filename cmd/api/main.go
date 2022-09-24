package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"

	"github.com/tiagocesar/yak-shop/handler"
	"github.com/tiagocesar/yak-shop/internal/models"
	"github.com/tiagocesar/yak-shop/internal/yak"
)

const defaultFilePath = "./herd-sample.xml"

func main() {
	filepath := flag.String("file", defaultFilePath, "The file path for the XML path with herd info")

	flag.Parse()

	file, err := os.ReadFile(*filepath)
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
