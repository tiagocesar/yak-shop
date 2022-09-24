package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type yakProcessor interface {
	Process(ctx context.Context, day int) (int, int, error)
}

type httpServer struct {
	yakService yakProcessor
}

func NewHttpServer(yakService yakProcessor) *httpServer {
	return &httpServer{yakService: yakService}
}

func (h *httpServer) ConfigureAndServe(port string) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	router.Get("/yak-shop/stock/{day}", nil)
	router.Get("/yak-shop/herd/{day}", nil)
	router.Post("/yak-shop/order/{day}", nil)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Failed to start HTTP server")
	}
}
