package valuereceiver

import "testing"

func TestCredit(t *testing.T) {
	w := &Wallet{Balance: 100}
	w.Credit(50)
	w.Credit(25)
	if w.Balance != 175 {
		t.Errorf("Balance=%d; want 175", w.Balance)
	}
}
