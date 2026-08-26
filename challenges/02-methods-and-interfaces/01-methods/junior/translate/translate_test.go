package translate

import "testing"

func TestTranslate(t *testing.T) {
	cases := []struct {
		name   string
		p      Point
		dx, dy float64
		want   Point
	}{
		{"positive", Point{1, 2}, 3, 4, Point{4, 6}},
		{"negative", Point{5, 5}, -2, -3, Point{3, 2}},
		{"zero", Point{7, 8}, 0, 0, Point{7, 8}},
		{"origin", Point{0, 0}, 1, 1, Point{1, 1}},
		{"decimal", Point{0.5, 0.5}, 0.5, 0.5, Point{1, 1}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := tc.p
			p.Translate(tc.dx, tc.dy)
			if p != tc.want {
				t.Errorf("Point{%g,%g}.Translate(%g,%g) => {%g,%g}, want {%g,%g}",
					tc.p.X, tc.p.Y, tc.dx, tc.dy, p.X, p.Y, tc.want.X, tc.want.Y)
			}
		})
	}
}
