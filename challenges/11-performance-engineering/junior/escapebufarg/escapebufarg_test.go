package escapebufarg

import "testing"

var sink Point

func TestAdd(t *testing.T) {
	if got := Add(Point{1, 2}, Point{3, 4}); got != (Point{4, 6}) {
		t.Errorf("Add = %v, want {4 6}", got)
	}
	if got := Add(Point{}, Point{}); got != (Point{}) {
		t.Errorf("Add = %v, want {0 0}", got)
	}
	if got := Add(Point{-1, 0}, Point{1, 0}); got != (Point{}) {
		t.Errorf("Add = %v, want {0 0}", got)
	}
}

func TestAddDoesNotModifyInputs(t *testing.T) {
	a := Point{1, 2}
	b := Point{3, 4}
	Add(a, b)
	if a != (Point{1, 2}) || b != (Point{3, 4}) {
		t.Errorf("inputs changed: a = %v, b = %v", a, b)
	}
}

func TestAddInto(t *testing.T) {
	var p Point
	AddInto(&p, Point{1, 2}, Point{3, 4})
	if p != (Point{4, 6}) {
		t.Errorf("p = %v, want {4 6}", p)
	}
	AddInto(&p, p, Point{1, 1})
	if p != (Point{5, 7}) {
		t.Errorf("p = %v, want {5 7}", p)
	}
}

func TestNeitherAllocates(t *testing.T) {
	if allocs := testing.AllocsPerRun(100, func() { sink = Add(Point{1, 2}, Point{3, 4}) }); allocs != 0 {
		t.Errorf("Add made %v allocations, want 0", allocs)
	}
	var p Point
	if allocs := testing.AllocsPerRun(100, func() { AddInto(&p, Point{1, 2}, Point{3, 4}) }); allocs != 0 {
		t.Errorf("AddInto made %v allocations, want 0", allocs)
	}
}
