package closurecapture

import "testing"

var sink func() int

func TestCounter(t *testing.T) {
	c := Counter(1)
	for want := 1; want <= 4; want++ {
		if got := c(); got != want {
			t.Fatalf("call = %d, want %d", got, want)
		}
	}
}

func TestCountersAreIndependent(t *testing.T) {
	a, b := Counter(10), Counter(10)
	a()
	a()
	if got := b(); got != 10 {
		t.Errorf("b() = %d, want 10: the counters share state", got)
	}
}

func TestCounterFromZeroAndNegative(t *testing.T) {
	c := Counter(-2)
	if got := []int{c(), c(), c()}; got[0] != -2 || got[1] != -1 || got[2] != 0 {
		t.Errorf("got %v, want [-2 -1 0]", got)
	}
}

func TestCounterAllocationsAreBounded(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { sink = Counter(0) }); n > 2 {
		t.Errorf("Counter made %v allocations, want at most 2", n)
	}
}
