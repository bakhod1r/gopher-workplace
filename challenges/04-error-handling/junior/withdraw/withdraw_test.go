package withdraw

import (
	"errors"
	"testing"
)

func TestWithdraw(t *testing.T) {
	cases := []struct {
		name            string
		balance, amount int
		want            int
		wantErr         error
	}{
		{"partial", 100, 30, 70, nil},
		{"exact", 100, 100, 0, nil},
		{"zero_amount", 100, 0, 100, ErrInvalidAmount},
		{"negative_amount", 100, -5, 100, ErrInvalidAmount},
		{"overdraft", 100, 150, 100, ErrInsufficientFunds},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Withdraw(tc.balance, tc.amount)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Withdraw(%d, %d) err = %v, want %v", tc.balance, tc.amount, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Withdraw(%d, %d) = %d, want %d", tc.balance, tc.amount, got, tc.want)
			}
		})
	}
}
