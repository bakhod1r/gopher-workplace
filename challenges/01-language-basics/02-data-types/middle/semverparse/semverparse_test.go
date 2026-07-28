package semverparse

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		s          string
		ma, mi, pa int
		ok         bool
	}{
		{"1.4.10", 1, 4, 10, true},
		{"0.0.0", 0, 0, 0, true},
		{"2.0", 0, 0, 0, false},
		{"1.2.3.4", 0, 0, 0, false},
		{"1.x.0", 0, 0, 0, false},
		{"", 0, 0, 0, false},
	}
	for _, c := range cases {
		ma, mi, pa, ok := Parse(c.s)
		if ma != c.ma || mi != c.mi || pa != c.pa || ok != c.ok {
			t.Errorf("Parse(%q)=(%d,%d,%d,%v); want (%d,%d,%d,%v)",
				c.s, ma, mi, pa, ok, c.ma, c.mi, c.pa, c.ok)
		}
	}
}
