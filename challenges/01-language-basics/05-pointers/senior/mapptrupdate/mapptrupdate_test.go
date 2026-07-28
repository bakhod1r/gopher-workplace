package mapptrupdate

import "testing"

func TestCredit(t *testing.T) {
	a := &Account{Balance: 100}
	m := map[int]*Account{1: a}
	if !Credit(m, 1, 50) {
		t.Fatalf("should succeed")
	}
	if a.Balance != 150 {
		t.Errorf("Balance=%d want 150 (mutated a copy)", a.Balance)
	}
}
