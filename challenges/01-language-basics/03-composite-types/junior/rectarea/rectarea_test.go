package rectarea

import "testing"

func TestArea(t *testing.T) {
	if got := (Rect{3, 4}).Area(); got != 12 {
		t.Errorf("Area=%d; want 12", got)
	}
}

func TestScale(t *testing.T) {
	r := Rect{2, 3}
	s := r.Scale(2)
	if s != (Rect{4, 6}) {
		t.Errorf("Scale=%v; want {4 6}", s)
	}
	if r != (Rect{2, 3}) {
		t.Error("original mutated")
	}
}
