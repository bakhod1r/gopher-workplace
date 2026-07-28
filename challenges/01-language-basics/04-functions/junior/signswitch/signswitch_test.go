package signswitch

import "testing"

func TestSign(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{-4, "negative"}, {0, "zero"}, {7, "positive"},
	}
	for _, c := range cases {
		if got := Sign(c.n); got != c.want {
			t.Errorf("Sign(%d)=%q want %q", c.n, got, c.want)
		}
	}
}
