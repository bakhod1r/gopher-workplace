package pipeline

import "testing"

func TestPipe(t *testing.T) {
	inc := func(x int) int { return x + 1 }
	dbl := func(x int) int { return x * 2 }
	if got := Pipe(3, inc, dbl); got != 8 {
		t.Errorf("=%d want 8 (dbl(inc(3)))", got)
	}
	if got := Pipe(5); got != 5 {
		t.Errorf("no fns should be identity: %d", got)
	}
}
