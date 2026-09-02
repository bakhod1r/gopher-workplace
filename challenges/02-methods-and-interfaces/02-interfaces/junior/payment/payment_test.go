package payment

import "testing"

func TestCheckout(t *testing.T) {
	cases := []struct {
		name   string
		p      Payment
		amount int
		want   string
	}{
		{"card_ok", Card{Limit: 100}, 50, "paid"},
		{"card_over", Card{Limit: 100}, 150, "declined"},
		{"card_exact", Card{Limit: 100}, 100, "paid"},
		{"cash_exact", Cash{Available: 20}, 20, "paid"},
		{"cash_over", Cash{Available: 20}, 21, "declined"},
		{"zero", Cash{Available: 0}, 0, "paid"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Checkout(tc.p, tc.amount); got != tc.want {
				t.Errorf("Checkout = %q, want %q", got, tc.want)
			}
		})
	}
}
