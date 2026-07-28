package tracker

import "testing"

func TestTracker(t *testing.T) {
	add, total := NewTracker()
	if total() != 0 {
		t.Errorf("start should be 0")
	}
	add(5)
	add(3)
	if total() != 8 {
		t.Errorf("=%d want 8", total())
	}
}
