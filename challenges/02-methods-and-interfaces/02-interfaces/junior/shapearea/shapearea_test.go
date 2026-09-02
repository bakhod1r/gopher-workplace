package shapearea

import "testing"

func TestArea(t *testing.T) {
	if got := (Rect{W: 3, H: 4}).Area(); got != 12 {
		t.Errorf("Rect.Area = %v, want 12", got)
	}
	if got := (Square{Side: 5}).Area(); got != 25 {
		t.Errorf("Square.Area = %v, want 25", got)
	}
	if got := (Rect{}).Area(); got != 0 {
		t.Errorf("zero Rect.Area = %v, want 0", got)
	}
}

func TestTotalArea(t *testing.T) {
	cases := []struct {
		name   string
		shapes []Shape
		want   float64
	}{
		{"mixed", []Shape{Rect{W: 2, H: 2}, Square{Side: 3}}, 13},
		{"empty", nil, 0},
		{"fractional", []Shape{Rect{W: 0.5, H: 4}}, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TotalArea(tc.shapes); got != tc.want {
				t.Errorf("TotalArea = %v, want %v", got, tc.want)
			}
		})
	}
}
