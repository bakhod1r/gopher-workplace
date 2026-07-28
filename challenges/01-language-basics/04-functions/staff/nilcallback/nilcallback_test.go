package nilcallback

import "testing"

func TestProcess(t *testing.T) {
	if got := Process([]int{1, 2, 3}, nil); got != 6 {
		t.Errorf("nil hook should sum unchanged: %d want 6", got)
	}
	if got := Process([]int{1, 2, 3}, func(x int) int { return x * 10 }); got != 60 {
		t.Errorf("=%d want 60", got)
	}
}
