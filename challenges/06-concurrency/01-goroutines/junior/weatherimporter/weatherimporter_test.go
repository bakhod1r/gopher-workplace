package weatherimporter

import (
	"reflect"
	"testing"
)

func TestToCelsius(t *testing.T) {
	cases := []struct {
		name     string
		readings []float64
		want     []float64
	}{
		{"freeze_and_boil", []float64{32, 212}, []float64{0, 100}},
		{"fixed_point", []float64{-40}, []float64{-40}},
		{"body_and_cold", []float64{104, 14}, []float64{40, -10}},
		{"empty", []float64{}, []float64{}},
		{"single", []float64{50}, []float64{10}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToCelsius(tc.readings); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ToCelsius(%v) = %v, want %v", tc.readings, got, tc.want)
			}
		})
	}
}
