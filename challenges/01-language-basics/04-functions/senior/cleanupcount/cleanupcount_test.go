package cleanupcount

import "testing"

func TestPeakThenDrain(t *testing.T) {
	if got := PeakThenDrain(3); got != 3 {
		t.Errorf("=%d want 3 (all three open before any defer drains)", got)
	}
	if got := PeakThenDrain(0); got != 0 {
		t.Errorf("=%d want 0", got)
	}
}
