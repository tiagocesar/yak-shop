package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"github.com/tiagocesar/yak-shop/internal/models"
)

type mockYakSvc struct{}

func (m *mockYakSvc) Process(_ int) ([]models.Yak, float32, int) {
	return nil, 1104.480, 3
}

func Test_orderHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       func() string
		statusCode int
	}{
		{
			name: "missing customer info",
			body: func() string {
				req := orderHandlerRequest{}

				j, _ := json.Marshal(req)

				return string(j)
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "missing order info",
			body: func() string {
				req := orderHandlerRequest{
					Customer: "Sample customer",
				}

				j, _ := json.Marshal(req)

				return string(j)
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "successful order",
			body: func() string {
				req := orderHandlerRequest{
					Customer: "Sample customer",
					Order: struct {
						Milk  float32 `json:"milk"`
						Skins int     `json:"skins"`
					}{
						Milk:  1100,
						Skins: 3,
					},
				}

				j, _ := json.Marshal(req)

				return string(j)
			},
			statusCode: http.StatusCreated,
		},
		{
			name: "successful order - partially fulfilled",
			body: func() string {
				req := orderHandlerRequest{
					Customer: "Sample customer",
					Order: struct {
						Milk  float32 `json:"milk"`
						Skins int     `json:"skins"`
					}{
						Milk:  1500,
						Skins: 3,
					},
				}

				j, _ := json.Marshal(req)

				return string(j)
			},
			statusCode: http.StatusPartialContent,
		},
		{
			name: "failed order - not enough goods",
			body: func() string {
				req := orderHandlerRequest{
					Customer: "Sample customer",
					Order: struct {
						Milk  float32 `json:"milk"`
						Skins int     `json:"skins"`
					}{
						Milk:  1500,
						Skins: 5,
					},
				}

				j, _ := json.Marshal(req)

				return string(j)
			},
			statusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rr := httptest.NewRecorder()

			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("day", "14")

			req, err := http.NewRequestWithContext(
				context.Background(), http.MethodPost, "/yak-shop/order/{day}", strings.NewReader(test.body()))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
			require.NoError(t, err)

			h := HttpServer{
				yakService: &mockYakSvc{},
			}

			h.orderHandler(rr, req)

			require.Equal(t, test.statusCode, rr.Code)
		})
	}
}
