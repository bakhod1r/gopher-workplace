package sharedcell

import "testing"

func TestCounters(t *testing.T) {
	cs := Counters(2)
	if cs[0]() != 1 || cs[0]() != 2 {
		t.Errorf("counter 0 not 1,2")
	}
	if cs[1]() != 1 {
		t.Errorf("counter 1 not independent (shared cell?)")
	}
}
