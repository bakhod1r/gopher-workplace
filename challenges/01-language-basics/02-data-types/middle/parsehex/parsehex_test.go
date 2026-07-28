package parsehex

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		s  string
		n  int
		ok bool
	}{
		{"ff", 255, true}, {"1A2B", 6699, true}, {"0", 0, true},
		{"", 0, false}, {"xy", 0, false}, {"1g", 0, false},
	}
	for _, c := range cases {
		n, ok := Parse(c.s)
		if n != c.n || ok != c.ok {
			t.Errorf("Parse(%q)=(%d,%v); want (%d,%v)", c.s, n, ok, c.n, c.ok)
		}
	}
}
