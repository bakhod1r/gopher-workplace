package runeat

import "testing"

func TestAt(t *testing.T) {
	cases := []struct {
		s  string
		n  int
		r  rune
		ok bool
	}{
		{"hello", 1, 'e', true},
		{"héllo", 2, 'l', true},
		{"日本", 1, '本', true},
		{"日本", 2, 0, false}, // 2 runes, index 2 out of range
		{"ab", -1, 0, false},
	}
	for _, c := range cases {
		r, ok := At(c.s, c.n)
		if r != c.r || ok != c.ok {
			t.Errorf("At(%q,%d)=(%q,%v); want (%q,%v)", c.s, c.n, r, ok, c.r, c.ok)
		}
	}
}
