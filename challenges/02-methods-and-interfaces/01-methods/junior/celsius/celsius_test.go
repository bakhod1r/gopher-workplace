package celsius

import (
	"math"
	"testing"
)

func TestToFahrenheit(t *testing.T) {
	cases := []struct {
		name string
		c    Celsius
		want float64
	}{
		{"freezing", 0, 32},
		{"boiling", 100, 212},
		{"body", 37, 98.6},
		{"negative", -40, -40},
		{"sub_zero", -10, 14},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.c.ToFahrenheit()
			if math.Abs(got-tc.want) > 0.001 {
				t.Errorf("Celsius(%g).ToFahrenheit() = %g, want %g",
					float64(tc.c), got, tc.want)
			}
		})
	}
}
