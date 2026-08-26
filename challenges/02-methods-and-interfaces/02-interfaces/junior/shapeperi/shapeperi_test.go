package shapeperi

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	c := Circle{Radius: 1}
	if got := c.Perimeter(); math.Abs(got-2*math.Pi) > 0.001 {
		t.Errorf("Circle.Perimeter() = %f", got)
	}

	s := Square{Side: 5}
	if got := s.Perimeter(); got != 20 {
		t.Errorf("Square.Perimeter() = %f", got)
	}
}
