package almostequal

import "testing"

func TestAlmostEqual(t *testing.T) {
	cases := []struct {
		name string
		a, b float64
		want bool
	}{
		{"float rounding", 0.1 + 0.2, 0.3, true},
		{"identical", 1.0, 1.0, true},
		{"clearly different", 1.0, 1.0001, false},
		{"negative equal", -2.5, -2.5, true},
		{"just outside tolerance", 0, 1e-6, false},
		{"just inside tolerance", 0, 1e-12, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := AlmostEqual(tc.a, tc.b); got != tc.want {
				t.Errorf("AlmostEqual(%g, %g) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
