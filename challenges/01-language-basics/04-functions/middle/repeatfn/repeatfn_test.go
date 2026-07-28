package repeatfn

import "testing"

func TestRepeat(t *testing.T) {
	inc := func(x int) int { return x + 1 }
	if got := Repeat(inc, 3)(0); got != 3 {
		t.Errorf("=%d want 3", got)
	}
	if got := Repeat(inc, 0)(9); got != 9 {
		t.Errorf("n=0 identity, got %d", got)
	}
}
