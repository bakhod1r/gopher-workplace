package scaleall

import "testing"

func TestScaleAll(t *testing.T) {
	a, b := 2, 3
	ScaleAll([]*int{&a, nil, &b}, 10)
	if a != 20 || b != 30 {
		t.Errorf("a,b=%d,%d want 20,30", a, b)
	}
}
