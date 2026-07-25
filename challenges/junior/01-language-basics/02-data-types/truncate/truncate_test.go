package truncate

import "testing"

func TestWholePart(t *testing.T) {
	cases := []struct {
		name string
		in   float64
		want int
	}{
		{"zero", 0, 0},
		{"positive fraction drops", 9.99, 9},
		{"exact integer", 4.0, 4},
		{"just under one", 0.999, 0},
		{"negative truncates toward zero", -9.99, -9},
		{"negative just under zero", -0.5, 0},
		{"large value", 12345.678, 12345},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WholePart(tc.in); got != tc.want {
				t.Errorf("WholePart(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
