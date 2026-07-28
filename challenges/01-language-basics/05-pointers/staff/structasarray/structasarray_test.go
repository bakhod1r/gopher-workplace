package structasarray

import "testing"

func TestSum(t *testing.T) {
	p := &Pair{A: 3, B: 4}
	if got := Sum(p); got != 7 {
		t.Errorf("=%d want 7 (only read the first field?)", got)
	}
}
