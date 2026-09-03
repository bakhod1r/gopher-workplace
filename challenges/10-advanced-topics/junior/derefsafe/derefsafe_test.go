package derefsafe

import "testing"

var sink int

func TestValue(t *testing.T) {
	n := 42
	if got := Value(&n); got != 42 {
		t.Errorf("Value = %d, want 42", got)
	}
	if got := Value(nil); got != 0 {
		t.Errorf("Value(nil) = %d, want 0", got)
	}
	zero := 0
	if got := Value(&zero); got != 0 {
		t.Errorf("Value = %d, want 0", got)
	}
}

func TestValueSeesLaterWrites(t *testing.T) {
	n := 1
	p := &n
	n = 7
	if got := Value(p); got != 7 {
		t.Errorf("Value = %d, want 7: the pointer is a live view", got)
	}
}

func TestValueAllocatesNothing(t *testing.T) {
	n := 3
	p := &n
	if got := testing.AllocsPerRun(200, func() { sink = Value(p) }); got != 0 {
		t.Errorf("Value made %v allocations, want 0", got)
	}
}
