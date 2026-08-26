package withdraw

import "testing"

func TestWithdraw(t *testing.T) {
	cases := []struct {
		name    string
		start   int
		amount  int
		wantOK  bool
		wantBal int
	}{
		{"normal", 100, 30, true, 70},
		{"insufficient", 100, 200, false, 100},
		{"exact", 100, 100, true, 0},
		{"negative", 100, -1, false, 100},
		{"zero", 100, 0, false, 100},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := Account{Balance: tc.start}
			ok := a.Withdraw(tc.amount)
			if ok != tc.wantOK {
				t.Errorf("Account{%d}.Withdraw(%d) ok = %v, want %v",
					tc.start, tc.amount, ok, tc.wantOK)
			}
			if a.Balance != tc.wantBal {
				t.Errorf("Account{%d}.Withdraw(%d) => Balance = %d, want %d",
					tc.start, tc.amount, a.Balance, tc.wantBal)
			}
		})
	}
}
