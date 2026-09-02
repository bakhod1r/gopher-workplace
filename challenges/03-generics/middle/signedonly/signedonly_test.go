package signedonly

import "testing"

func TestNegate(t *testing.T) {
	if got := Negate(3); got != -3 {
		t.Errorf("Negate(3) = %v, want -3", got)
	}
	if got := Negate(-3); got != 3 {
		t.Errorf("Negate(-3) = %v, want 3", got)
	}
	if got := Negate(int64(0)); got != 0 {
		t.Errorf("Negate(0) = %v, want 0", got)
	}
}

func TestAbsDiff(t *testing.T) {
	if got := AbsDiff(2, 5); got != 3 {
		t.Errorf("AbsDiff(2, 5) = %v, want 3", got)
	}
	if got := AbsDiff(5, 2); got != 3 {
		t.Errorf("AbsDiff(5, 2) = %v, want 3", got)
	}
	if got := AbsDiff(4, 4); got != 0 {
		t.Errorf("AbsDiff(4, 4) = %v, want 0", got)
	}
	if got := AbsDiff(int64(-3), int64(3)); got != 6 {
		t.Errorf("AbsDiff(-3, 3) = %v, want 6", got)
	}
}
