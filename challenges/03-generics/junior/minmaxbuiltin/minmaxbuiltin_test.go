package minmaxbuiltin

import "testing"

func TestMiddle(t *testing.T) {
	if got := Middle(5, 0, 3); got != 3 {
		t.Errorf("Middle(5, 0, 3) = %v, want 3", got)
	}
	if got := Middle(-1, 0, 3); got != 0 {
		t.Errorf("Middle(-1, 0, 3) = %v, want 0", got)
	}
	if got := Middle(2, 0, 3); got != 2 {
		t.Errorf("Middle(2, 0, 3) = %v, want 2", got)
	}
	if got := Middle(2.5, 0.0, 1.0); got != 1.0 {
		t.Errorf("Middle(2.5, 0, 1) = %v, want 1", got)
	}
}

func TestSpread(t *testing.T) {
	if got := Spread(1, 9, 4); got != 8 {
		t.Errorf("Spread(1, 9, 4) = %v, want 8", got)
	}
	if got := Spread(3, 3, 3); got != 0 {
		t.Errorf("Spread(3, 3, 3) = %v, want 0", got)
	}
	if got := Spread(-5, 0, 5); got != 10 {
		t.Errorf("Spread(-5, 0, 5) = %v, want 10", got)
	}
}
