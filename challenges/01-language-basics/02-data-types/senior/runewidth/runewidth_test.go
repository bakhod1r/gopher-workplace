package runewidth

import "testing"

func TestFirstWidth(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"a", 1},
		{"é", 2},
		{"日本", 3}, // first rune only
		{"", 0},
	}
	for _, c := range cases {
		if got := FirstWidth(c.s); got != c.want {
			t.Errorf("FirstWidth(%q)=%d; want %d", c.s, got, c.want)
		}
	}
}
