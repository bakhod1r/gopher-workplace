package spreadmissing

import "testing"

func TestTotal(t *testing.T) {
	if got := Total([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("=%d want 10", got)
	}
	if got := Total(nil); got != 0 {
		t.Errorf("=%d want 0", got)
	}
}
