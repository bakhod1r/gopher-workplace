package durationparse

import "testing"

func TestSeconds(t *testing.T) {
	cases := []struct {
		s  string
		n  int
		ok bool
	}{
		{"45s", 45, true},
		{"2m", 120, true},
		{"1h30m", 5400, true},
		{"1h30m45s", 5445, true},
		{"10", 0, false},
		{"", 0, true},
	}
	for _, c := range cases {
		n, ok := Seconds(c.s)
		if n != c.n || ok != c.ok {
			t.Errorf("Seconds(%q)=(%d,%v); want (%d,%v)", c.s, n, ok, c.n, c.ok)
		}
	}
}
