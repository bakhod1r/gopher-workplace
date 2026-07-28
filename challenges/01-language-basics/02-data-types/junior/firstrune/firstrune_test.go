package firstrune

import "testing"

func TestFirst(t *testing.T) {
	cases := []struct {
		s    string
		want rune
	}{
		{"hello", 'h'},
		{"étage", 'é'},
		{"日本", '日'},
		{"", 0},
	}
	for _, c := range cases {
		if got := First(c.s); got != c.want {
			t.Errorf("First(%q)=%q(%d); want %q(%d)", c.s, got, got, c.want, c.want)
		}
	}
}
