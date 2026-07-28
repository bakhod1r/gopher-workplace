package absif

import "testing"

func TestAbs(t *testing.T) {
	cases := []struct{ n, want int }{{-5, 5}, {5, 5}, {0, 0}}
	for _, c := range cases {
		if got := Abs(c.n); got != c.want {
			t.Errorf("Abs(%d)=%d want %d", c.n, got, c.want)
		}
	}
}
