package yak

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tiagocesar/yak-shop/internal/models"
)

func Test_Process(t *testing.T) {
	t.Run("Should process the herd for a new day", func(t *testing.T) {
		t.Parallel()

		herd := []models.YakImport{
			{Age: 4},
			{Age: 8},
			{Age: 9.5},
		}

		svc := NewService(herd)

		_, totalMilk, totalWool := svc.Process(13)

		require.Equal(t, float32(1104.48), totalMilk)
		require.Equal(t, 3, totalWool)
	})
}
