package sumptrs

import "testing"

func TestSumPtrs(t *testing.T) {
	a, b, c := 1, 2, 3
	if got := SumPtrs([]*int{&a, nil, &b, &c}); got != 6 {
		t.Errorf("=%d want 6", got)
	}
}
