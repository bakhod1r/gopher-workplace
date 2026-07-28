package utf8count

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"hello", 5},
		{"héllo", 5},
		{"日本語", 3},
		{"a€b", 3},
		{"", 0},
	}
	for _, c := range cases {
		if got := Count([]byte(c.s)); got != c.want {
			t.Errorf("Count(%q)=%d; want %d", c.s, got, c.want)
		}
	}
}
