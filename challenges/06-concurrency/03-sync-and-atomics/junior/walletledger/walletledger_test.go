package walletledger

import (
	"sync"
	"testing"
)

func TestWallet(t *testing.T) {
	cases := []struct {
		name    string
		opening int
		credit  int
		debit   int
		wantOK  bool
		want    int
	}{
		{"top_up_only", 100, 50, 0, true, 150},
		{"charge_ok", 100, 0, 30, true, 70},
		{"charge_exact", 50, 0, 50, true, 0},
		{"insufficient_funds", 20, 0, 50, false, 20},
		{"top_up_then_charge", 0, 80, 30, true, 50},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := NewWallet(tc.opening)
			if tc.credit != 0 {
				w.Credit(tc.credit)
			}
			if got := w.Debit(tc.debit); got != tc.wantOK {
				t.Errorf("Debit(%d) = %v, want %v", tc.debit, got, tc.wantOK)
			}
			if got := w.Balance(); got != tc.want {
				t.Errorf("Balance() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestWalletConcurrent(t *testing.T) {
	w := NewWallet(1000)
	const charges = 20
	ok := make(chan bool, charges)
	var wg sync.WaitGroup
	wg.Add(charges)
	for i := 0; i < charges; i++ {
		go func() {
			defer wg.Done()
			ok <- w.Debit(100)
			w.Balance()
		}()
	}
	wg.Wait()
	close(ok)

	granted := 0
	for v := range ok {
		if v {
			granted++
		}
	}
	if granted != 10 {
		t.Errorf("successful charges = %d, want 10", granted)
	}
	if got := w.Balance(); got != 0 {
		t.Errorf("Balance() = %d, want 0", got)
	}
}
