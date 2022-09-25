package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_orderHandlerRequest_validate(t *testing.T) {
	tests := []struct {
		name        string
		req         orderHandlerRequest
		expectedErr error
	}{
		{
			name: "success",
			req: orderHandlerRequest{
				Customer: "Sample customer",
				Order: struct {
					Milk  float32 `json:"milk"`
					Skins int     `json:"skins"`
				}{
					Milk:  1100,
					Skins: 3,
				},
			},
		},
		{
			name:        "error - missing customer information",
			req:         orderHandlerRequest{},
			expectedErr: ErrCustomerNotSpecified,
		},
		{
			name: "error - missing order goods",
			req: orderHandlerRequest{
				Customer: "Sample customer",
			},
			expectedErr: ErrNoGoodsSpecified,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			err := test.req.validate()

			require.Equal(t, err, test.expectedErr)
		})
	}
}
