package byvalue

import "testing"

var sink Point

func TestScale(t *testing.T) {
	if got := Scale(Point{2, 3}, 2); got != (Point{4, 6}) {
		t.Errorf("Scale = %v, want {4 6}", got)
	}
	if got := Scale(Point{1, 1}, 0); got != (Point{0, 0}) {
		t.Errorf("Scale = %v, want {0 0}", got)
	}
	if got := Scale(Point{-1, 2}, 3); got != (Point{-3, 6}) {
		t.Errorf("Scale = %v, want {-3 6}", got)
	}
}

func TestScaleDoesNotTouchTheCaller(t *testing.T) {
	p := Point{2, 3}
	Scale(p, 5)
	if p != (Point{2, 3}) {
		t.Errorf("p = %v, want {2 3}: the parameter is a copy", p)
	}
}

func TestScaleAllocatesNothing(t *testing.T) {
	p := Point{2, 3}
	if n := testing.AllocsPerRun(100, func() { sink = Scale(p, 2) }); n != 0 {
		t.Errorf("Scale made %v allocations, want 0", n)
	}
}
