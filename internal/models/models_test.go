package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Milk(t *testing.T) {
	t.Run("Should calculate the correct amount of milk", func(t *testing.T) {
		yak := Yak{
			AgeInDays: 100,
		}

		res := yak.Milk()

		require.Equal(t, float32(47), res)
	})
}

func Test_Shave(t *testing.T) {
	tests := []struct {
		name     string
		yak      Yak
		expected int
	}{
		{
			name: "yak should be shaven",
			yak: Yak{
				AgeInDays: 400,
			},
			expected: 1,
		},
		{
			name: "yak is dead so should not be shaven",
			yak: Yak{
				Dead: true,
			},
		},
		{
			name: "yak is not ready to be shaven",
			yak: Yak{
				AgeInDays: 400,
				NextShave: 11,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			res := test.yak.Shave(10)

			require.Equal(t, test.expected, res)
		})
	}
}

func Test_Age(t *testing.T) {
	t.Run("Yak should age and be alive", func(t *testing.T) {
		t.Parallel()

		yak := Yak{AgeInDays: 400}

		yak.Age()

		require.Equal(t, 401, yak.AgeInDays)
	})

	t.Run("Yak should age and die", func(t *testing.T) {
		t.Parallel()

		yak := Yak{AgeInDays: 999}

		yak.Age()

		require.True(t, yak.Dead)
	})
}
