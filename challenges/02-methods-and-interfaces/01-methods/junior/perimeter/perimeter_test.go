package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	cases := []struct {
		name string
		r    Rect
		want float64
	}{
		{"3x4", Rect{3, 4}, 14},
		{"zero_width", Rect{0, 5}, 10},
		{"zero_both", Rect{0, 0}, 0},
		{"unit", Rect{1, 1}, 4},
		{"decimal", Rect{2.5, 3.5}, 12},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.r.Perimeter(); got != tc.want {
				t.Errorf("Rect{%g, %g}.Perimeter() = %g, want %g",
					tc.r.Width, tc.r.Height, got, tc.want)
			}
		})
	}
}
