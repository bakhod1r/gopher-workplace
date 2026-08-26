package area

import "testing"

func TestArea(t *testing.T) {
	cases := []struct {
		name string
		r    Rect
		want float64
	}{
		{"3x4", Rect{3, 4}, 12},
		{"zero_width", Rect{0, 5}, 0},
		{"zero_height", Rect{7, 0}, 0},
		{"unit", Rect{1, 1}, 1},
		{"decimal", Rect{2.5, 4}, 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.r.Area(); got != tc.want {
				t.Errorf("Rect{%g, %g}.Area() = %g, want %g",
					tc.r.Width, tc.r.Height, got, tc.want)
			}
		})
	}
}
