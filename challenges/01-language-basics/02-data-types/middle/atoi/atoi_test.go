package atoi

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		s  string
		n  int
		ok bool
	}{
		{"0", 0, true}, {"42", 42, true}, {"-17", -17, true},
		{"", 0, false}, {"1a", 0, false}, {"-", 0, false},
	}
	for _, c := range cases {
		n, ok := Parse(c.s)
		if n != c.n || ok != c.ok {
			t.Errorf("Parse(%q)=(%d,%v); want (%d,%v)", c.s, n, ok, c.n, c.ok)
		}
	}
}
