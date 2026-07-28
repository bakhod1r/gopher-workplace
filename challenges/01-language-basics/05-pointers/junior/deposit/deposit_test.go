package deposit

import "testing"

func TestDeposit(t *testing.T) {
	a := Account{Balance: 100}
	Deposit(&a, 50)
	Deposit(&a, -30)
	if a.Balance != 120 {
		t.Errorf("Balance=%d want 120", a.Balance)
	}
}
