package stalecap

import "testing"

func TestFirstAfterGrow(t *testing.T) {
	if got := FirstAfterGrow(7); got != 99 {
		t.Errorf("=%d want 99 (write must hit the reallocated slice)", got)
	}
}
