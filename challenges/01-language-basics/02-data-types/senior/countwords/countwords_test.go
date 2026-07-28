package countwords

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"hello world", 2},
		{"  a   b  ", 2},
		{"one", 1},
		{"", 0},
		{"   ", 0},
	}
	for _, c := range cases {
		if got := Count(c.s); got != c.want {
			t.Errorf("Count(%q)=%d; want %d", c.s, got, c.want)
		}
	}
}
