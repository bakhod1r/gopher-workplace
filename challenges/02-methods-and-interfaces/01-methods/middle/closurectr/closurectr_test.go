package closurectr

import "testing"

func TestNewCounter(t *testing.T) {
	next := NewCounter()

	for i := 1; i <= 5; i++ {
		if got := next(); got != i {
			t.Errorf("call %d: got %d, want %d", i, got, i)
		}
	}

	// A second counter is independent.
	next2 := NewCounter()
	if got := next2(); got != 1 {
		t.Errorf("second counter first call: got %d, want 1", got)
	}
	// Original counter continues from 5.
	if got := next(); got != 6 {
		t.Errorf("first counter after 5: got %d, want 6", got)
	}
}
