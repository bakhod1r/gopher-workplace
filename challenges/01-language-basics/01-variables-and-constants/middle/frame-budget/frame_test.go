package frame

import "testing"

func TestBudget(t *testing.T) {
	if got := FrameBudgetMicros(); got != 16666 {
		t.Fatalf("budget=%d; want 16666", got)
	}
}

func TestOver(t *testing.T) {
	if OverBudget(10000) {
		t.Error("10000us is under budget")
	}
	if !OverBudget(20000) {
		t.Error("20000us is over budget")
	}
	if OverBudget(FrameBudgetMicros()) {
		t.Error("exactly at budget is not over")
	}
}
