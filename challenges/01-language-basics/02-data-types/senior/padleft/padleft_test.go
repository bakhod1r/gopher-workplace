package padleft

import "testing"

func TestPad(t *testing.T) {
	cases := []struct {
		s     string
		width int
		fill  byte
		want  string
	}{
		{"42", 5, '0', "00042"},
		{"abc", 3, ' ', "abc"},
		{"x", 4, '.', "...x"},
		{"toolong", 3, '0', "toolong"},
	}
	for _, c := range cases {
		if got := Pad(c.s, c.width, c.fill); got != c.want {
			t.Errorf("Pad(%q,%d,%q)=%q; want %q", c.s, c.width, c.fill, got, c.want)
		}
	}
}
