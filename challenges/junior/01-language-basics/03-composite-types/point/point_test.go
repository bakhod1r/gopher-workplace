package point

import "testing"

func TestTranslate(t *testing.T) {
	cases := []struct {
		name   string
		p      Point
		dx, dy int
		want   Point
	}{
		{"basic", Point{1, 2}, 3, 4, Point{4, 6}},
		{"negative", Point{0, 0}, -1, -1, Point{-1, -1}},
		{"zero shift", Point{5, 5}, 0, 0, Point{5, 5}},
		{"asymmetric", Point{2, 2}, 3, 7, Point{5, 9}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Translate(tc.p, tc.dx, tc.dy); got != tc.want {
				t.Errorf("Translate(%v, %d, %d) = %v, want %v", tc.p, tc.dx, tc.dy, got, tc.want)
			}
		})
	}
}

func TestTranslateDoesNotMutateCaller(t *testing.T) {
	p := Point{1, 2}
	_ = Translate(p, 10, 20)
	if p != (Point{1, 2}) {
		t.Errorf("caller's point was mutated: got %v, want {1 2}", p)
	}
}
