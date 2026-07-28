package money

import "testing"

func TestCents(t *testing.T) {
	cases := []struct {
		d    float64
		want int64
	}{
		{1.00, 100},
		{2.50, 250},
		{0.99, 99},
		{10.05, 1005},
	}
	for _, c := range cases {
		if got := Cents(c.d); got != c.want {
			t.Errorf("Cents(%v)=%d; want %d", c.d, got, c.want)
		}
	}
}
