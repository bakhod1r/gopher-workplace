package constraintifcmethod

import "testing"

func TestTotalCents(t *testing.T) {
	if got := TotalCents([]Book{{200}, {350}}); got != 550 {
		t.Errorf("TotalCents(books) = %d, want 550", got)
	}
	if got := TotalCents([]Coffee{{450}}); got != 450 {
		t.Errorf("TotalCents(coffees) = %d, want 450", got)
	}
	if got := TotalCents([]Book{}); got != 0 {
		t.Errorf("TotalCents([]) = %d, want 0", got)
	}
}
