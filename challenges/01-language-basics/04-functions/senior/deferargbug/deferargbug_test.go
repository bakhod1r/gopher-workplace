package deferargbug

import "testing"

func TestFinalCount(t *testing.T) {
	if got := FinalCount(5); got != 5 {
		t.Errorf("FinalCount(5)=%d want 5", got)
	}
	if got := FinalCount(0); got != 0 {
		t.Errorf("FinalCount(0)=%d want 0", got)
	}
}
