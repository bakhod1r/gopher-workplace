package shapearea

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	if got := c.Area(); math.Abs(got-math.Pi*25) > 0.001 {
		t.Errorf("Circle.Area() = %f", got)
	}
	if got := r.Area(); got != 12 {
		t.Errorf("Rectangle.Area() = %f", got)
	}

	total := TotalArea([]Shape{c, r})
	want := math.Pi*25 + 12
	if math.Abs(total-want) > 0.001 {
		t.Errorf("TotalArea = %f, want %f", total, want)
	}
}
