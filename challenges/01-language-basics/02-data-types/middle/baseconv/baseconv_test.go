package baseconv

import "testing"

func TestFormat(t *testing.T) {
	cases := []struct {
		n, base int
		want    string
	}{
		{0, 2, "0"}, {5, 2, "101"}, {255, 16, "ff"}, {10, 10, "10"}, {8, 8, "10"},
	}
	for _, c := range cases {
		if got := Format(c.n, c.base); got != c.want {
			t.Errorf("Format(%d,%d)=%q; want %q", c.n, c.base, got, c.want)
		}
	}
}
