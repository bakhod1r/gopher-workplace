package tierbug

import "testing"

func TestLabel(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{{1, "low"}, {2, "mid"}, {3, "high"}, {9, "?"}}
	for _, c := range cases {
		if got := Label(c.n); got != c.want {
			t.Errorf("Label(%d)=%q want %q", c.n, got, c.want)
		}
	}
}
