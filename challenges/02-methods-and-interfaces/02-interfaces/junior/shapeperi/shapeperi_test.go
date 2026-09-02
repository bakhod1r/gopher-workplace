package shapeperi

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	if got := (Rect{W: 2, H: 3}).Perimeter(); got != 10 {
		t.Errorf("Rect.Perimeter = %v, want 10", got)
	}
	if got := (Square{Side: 4}).Perimeter(); got != 16 {
		t.Errorf("Square.Perimeter = %v, want 16", got)
	}
	if got := (Circle{R: 1}).Perimeter(); math.Abs(got-2*math.Pi) > 1e-9 {
		t.Errorf("Circle.Perimeter = %v, want %v", got, 2*math.Pi)
	}
}

func TestLongestPerimeter(t *testing.T) {
	cases := []struct {
		name   string
		shapes []Shape
		want   float64
	}{
		{"rect_wins", []Shape{Square{Side: 1}, Rect{W: 5, H: 5}}, 20},
		{"square_wins", []Shape{Square{Side: 10}, Rect{W: 1, H: 1}}, 40},
		{"empty", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongestPerimeter(tc.shapes); got != tc.want {
				t.Errorf("LongestPerimeter = %v, want %v", got, tc.want)
			}
		})
	}
}
