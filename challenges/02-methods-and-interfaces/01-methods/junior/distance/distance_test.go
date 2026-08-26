package distance

import (
	"math"
	"testing"
)

func TestDistanceTo(t *testing.T) {
	cases := []struct {
		name string
		a, b Point
		want float64
	}{
		{"3-4-5", Point{0, 0}, Point{3, 4}, 5},
		{"same", Point{1, 1}, Point{1, 1}, 0},
		{"negative", Point{-1, -1}, Point{2, 3}, 5},
		{"horizontal", Point{0, 0}, Point{7, 0}, 7},
		{"vertical", Point{0, 0}, Point{0, 3}, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a.DistanceTo(tc.b)
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Point{%g,%g}.DistanceTo(Point{%g,%g}) = %g, want %g",
					tc.a.X, tc.a.Y, tc.b.X, tc.b.Y, got, tc.want)
			}
		})
	}
}
