package ifaceallocount

import "testing"

var sink int

func TestArea(t *testing.T) {
	if got := (&Rect{2, 3}).Area(); got != 6 {
		t.Errorf("Area = %d, want 6", got)
	}
	if got := (&Rect{}).Area(); got != 0 {
		t.Errorf("Area = %d, want 0", got)
	}
}

func TestPointerSatisfiesShape(t *testing.T) {
	var s Shape = &Rect{4, 5}
	if got := s.Area(); got != 20 {
		t.Errorf("Area through the interface = %d, want 20", got)
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{&Rect{2, 3}, &Rect{1, 1}, &Rect{0, 9}}
	if got := TotalArea(shapes); got != 7 {
		t.Errorf("TotalArea = %d, want 7", got)
	}
	if got := TotalArea(nil); got != 0 {
		t.Errorf("TotalArea(nil) = %d, want 0", got)
	}
}

func TestTotalAreaDoesNotAllocate(t *testing.T) {
	shapes := make([]Shape, 0, 100)
	for i := 0; i < 100; i++ {
		shapes = append(shapes, &Rect{i, i})
	}
	allocs := testing.AllocsPerRun(50, func() { sink = TotalArea(shapes) })
	if allocs != 0 {
		t.Errorf("TotalArea made %v allocations, want 0 — the shapes are already boxed", allocs)
	}
}
