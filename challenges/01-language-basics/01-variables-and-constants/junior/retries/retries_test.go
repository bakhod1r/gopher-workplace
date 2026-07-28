package retries

import "testing"

func TestMaxRetries(t *testing.T) {
	if MaxRetries != 3 {
		t.Errorf("MaxRetries = %d, want 3", MaxRetries)
	}
}

func TestBudget(t *testing.T) {
	if got := Budget(); got != 4 {
		t.Errorf("Budget() = %d, want 4", got)
	}
}
