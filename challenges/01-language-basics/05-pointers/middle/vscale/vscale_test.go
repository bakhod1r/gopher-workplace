package vscale

import "testing"

func TestScaled(t *testing.T) {
	p := Point{X: 1, Y: 2}
	q := p.Scaled(3)
	if q.X != 3 || q.Y != 6 {
		t.Errorf("q=%+v want {3 6}", q)
	}
	if p.X != 1 || p.Y != 2 {
		t.Errorf("receiver mutated: %+v", p)
	}
}
