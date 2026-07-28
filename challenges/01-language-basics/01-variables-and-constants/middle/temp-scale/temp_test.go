package temp

import "testing"

func TestConsts(t *testing.T) {
	if AbsoluteZero != -273.15 || Boiling != 100 {
		t.Fatalf("consts = %v,%v", AbsoluteZero, Boiling)
	}
}

func TestValid(t *testing.T) {
	cases := []struct {
		c  Celsius
		ok bool
	}{
		{-300, false}, {AbsoluteZero, true}, {0, true}, {Boiling, true},
	}
	for _, tc := range cases {
		if got := Valid(tc.c); got != tc.ok {
			t.Errorf("Valid(%v)=%v; want %v", tc.c, got, tc.ok)
		}
	}
}
