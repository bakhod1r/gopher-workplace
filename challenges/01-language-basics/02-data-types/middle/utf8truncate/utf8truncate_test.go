package utf8truncate

import "testing"

func TestTruncate(t *testing.T) {
	cases := []struct {
		s    string
		max  int
		want string
	}{
		{"hello", 3, "hel"},
		{"héllo", 2, "h"}, // é is 2 bytes; can't include half
		{"héllo", 3, "hé"},
		{"日本", 4, "日"}, // each kanji 3 bytes
		{"abc", 10, "abc"},
		{"abc", 0, ""},
	}
	for _, c := range cases {
		if got := Truncate(c.s, c.max); got != c.want {
			t.Errorf("Truncate(%q,%d)=%q; want %q", c.s, c.max, got, c.want)
		}
	}
}
