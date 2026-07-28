package base64val

import "testing"

func TestValue(t *testing.T) {
	cases := []struct {
		c  byte
		v  int
		ok bool
	}{
		{'A', 0, true}, {'Z', 25, true}, {'a', 26, true}, {'z', 51, true},
		{'0', 52, true}, {'9', 61, true}, {'+', 62, true}, {'/', 63, true},
		{'=', 0, false}, {' ', 0, false},
	}
	for _, c := range cases {
		v, ok := Value(c.c)
		if v != c.v || ok != c.ok {
			t.Errorf("Value(%q)=(%d,%v); want (%d,%v)", c.c, v, ok, c.v, c.ok)
		}
	}
}
