package runecount

import "testing"

func TestCharCount(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{{"abc", 3}, {"héllo", 5}, {"日本語", 3}, {"", 0}}
	for _, c := range cases {
		if got := CharCount(c.s); got != c.want {
			t.Errorf("CharCount(%q)=%d want %d", c.s, got, c.want)
		}
	}
}
