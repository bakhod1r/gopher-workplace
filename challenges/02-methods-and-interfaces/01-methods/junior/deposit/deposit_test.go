package deposit

import "testing"

func TestDeposit(t *testing.T) {
	cases := []struct {
		name    string
		start   int
		amount  int
		wantBal int
	}{
		{"normal", 100, 50, 150},
		{"negative_ignored", 100, -1, 100},
		{"zero_ignored", 100, 0, 100},
		{"from_zero", 0, 200, 200},
		{"large", 1000, 5000, 6000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := Account{Balance: tc.start}
			a.Deposit(tc.amount)
			if a.Balance != tc.wantBal {
				t.Errorf("Account{%d}.Deposit(%d) => Balance = %d, want %d",
					tc.start, tc.amount, a.Balance, tc.wantBal)
			}
		})
	}
}
