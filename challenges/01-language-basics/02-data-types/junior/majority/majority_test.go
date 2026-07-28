package majority

import "testing"

func TestMajority(t *testing.T) {
	cases := []struct {
		a, b, c, want bool
	}{
		{false, false, false, false},
		{true, false, false, false},
		{true, true, false, true},
		{false, true, true, true},
		{true, true, true, true},
	}
	for _, c := range cases {
		if got := Majority(c.a, c.b, c.c); got != c.want {
			t.Errorf("Majority(%v,%v,%v)=%v; want %v", c.a, c.b, c.c, got, c.want)
		}
	}
}
