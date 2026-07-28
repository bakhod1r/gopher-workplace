package countword

import "testing"

func TestCountEqual(t *testing.T) {
	if CountEqual(nil, 1) != 0 {
		t.Errorf("nil should be 0")
	}
	if got := CountEqual([]int{1, 2, 1, 3, 1}, 1); got != 3 {
		t.Errorf("=%d want 3", got)
	}
}
