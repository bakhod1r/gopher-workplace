package structupdate

import "testing"

func TestDeposit(t *testing.T) {
	a := &Account{Balance: 100}
	a.Deposit(50)
	a.Deposit(25)
	if a.Balance != 175 {
		t.Errorf("Balance=%d; want 175", a.Balance)
	}
}
