package moneyparse

import "testing"

func TestCents(t *testing.T) {
	cases := []struct {
		s  string
		c  int
		ok bool
	}{
		{"12.34", 1234, true},
		{"0.05", 5, true},
		{"7", 700, true},
		{"3.5", 350, true},
		{"1.234", 0, false},
		{"abc", 0, false},
		{"", 0, false},
	}
	for _, tc := range cases {
		c, ok := Cents(tc.s)
		if c != tc.c || ok != tc.ok {
			t.Errorf("Cents(%q)=(%d,%v); want (%d,%v)", tc.s, c, ok, tc.c, tc.ok)
		}
	}
}
