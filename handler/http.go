package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/tiagocesar/yak-shop/internal/models"
)

type yakProcessor interface {
	Process(day int) (float32, int)
	GetHerdInfo(day int) []models.Yak
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

	router.Get("/yak-shop/stock/{day}", h.stockHandler)
	router.Get("/yak-shop/herd/{day}", h.herdHandler)
	router.Post("/yak-shop/order/{day}", nil)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Failed to start HTTP server")
	}
}

// stockHandler returns a stock representation for the specified day in the request
func (h *httpServer) stockHandler(w http.ResponseWriter, r *http.Request) {
	d := chi.URLParam(r, "day")

	day, err := strconv.ParseInt(d, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(("Invalid day")))
		return
	}

	totalMilk, totalWool := h.yakService.Process(int(day))

	response := &stockHandlerResponse{
		Milk:  totalMilk,
		Skins: totalWool,
	}

	j, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(j)
}

// herdHandler returns information about the herd
func (h *httpServer) herdHandler(w http.ResponseWriter, r *http.Request) {
	d := chi.URLParam(r, "day")

	day, err := strconv.ParseInt(d, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(("Invalid day")))
		return
	}

	response := h.yakService.GetHerdInfo(int(day))
	herd := toHerdHandlerResponse(response)

	j, _ := json.Marshal(herd)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(j)
}
