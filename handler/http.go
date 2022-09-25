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
	Process(day int) ([]models.Yak, float32, int)
}

type HttpServer struct {
	yakService yakProcessor
}

func NewHttpServer(yakService yakProcessor) *HttpServer {
	return &HttpServer{yakService: yakService}
}

func (h *HttpServer) ConfigureAndServe(port string) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	router.Get("/yak-shop/stock/{day}", h.stockHandler)
	router.Get("/yak-shop/herd/{day}", h.herdHandler)
	router.Post("/yak-shop/order/{day}", h.orderHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Failed to start HTTP server")
	}
}

// stockHandler returns a stock representation for the specified day in the request
func (h *HttpServer) stockHandler(w http.ResponseWriter, r *http.Request) {
	d := chi.URLParam(r, "day")

	day, err := strconv.ParseInt(d, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(("Invalid day")))
		return
	}

	_, totalMilk, totalWool := h.yakService.Process(int(day))

	response := &stockHandlerResponse{
		Milk:  totalMilk,
		Skins: totalWool,
	}

	j, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(j)
}

// herdHandler returns information about the herd
func (h *HttpServer) herdHandler(w http.ResponseWriter, r *http.Request) {
	d := chi.URLParam(r, "day")

	day, err := strconv.ParseInt(d, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(("Invalid day")))
		return
	}

	response, _, _ := h.yakService.Process(int(day))
	herd := toHerdHandlerResponse(response)

	j, _ := json.Marshal(herd)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(j)
}

// orderHandler allows customers to place orders against our available stock.
// Orders are only processed if the entirety of each product is available
// according to the quantities specified by the customer.
// Partial order processing is possible.
func (h *HttpServer) orderHandler(w http.ResponseWriter, r *http.Request) {
	d := chi.URLParam(r, "day")

	day, err := strconv.ParseInt(d, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(("Invalid day")))
		return
	}

	// Parsing the request body
	var orderRequest orderHandlerRequest
	err = json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validating the request
	if err = orderRequest.validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, totalMilk, totalWool := h.yakService.Process(int(day))

	// Establishing what status code to return:
	//
	// 404 - there's no quantity to fulfill the order of either goods
	// 206 - there's one or another good in a quantity that fulfills part of the order
	// 201 - there's enough goods to fulfill the entire order
	statusCode := http.StatusNotFound
	var orderedMilk float32
	var orderedSkins int

	if orderRequest.Order.Milk <= totalMilk {
		statusCode = http.StatusPartialContent
		orderedMilk = orderRequest.Order.Milk
	}
	if orderRequest.Order.Skins <= totalWool {
		statusCode = http.StatusPartialContent
		orderedSkins = orderRequest.Order.Skins
	}

	// Checking if orders are fulfilled
	if orderedMilk > 0 && orderedSkins > 0 {
		statusCode = http.StatusCreated
	}

	orderResponse := orderHandlerResponse{
		Milk:  orderedMilk,
		Skins: orderedSkins,
	}

	j, _ := json.Marshal(orderResponse)

	w.WriteHeader(statusCode)
	_, _ = w.Write(j)
}
