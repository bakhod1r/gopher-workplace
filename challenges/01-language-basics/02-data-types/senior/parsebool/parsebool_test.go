package parsebool

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		s     string
		v, ok bool
	}{
		{"TRUE", true, true},
		{"on", true, true},
		{" no ", false, true},
		{"off", false, true}, // must be recognized as false
		{"maybe", false, false},
	}
	for _, c := range cases {
		v, ok := Parse(c.s)
		if v != c.v || ok != c.ok {
			t.Errorf("Parse(%q)=(%v,%v); want (%v,%v)", c.s, v, ok, c.v, c.ok)
		}
	}
}
