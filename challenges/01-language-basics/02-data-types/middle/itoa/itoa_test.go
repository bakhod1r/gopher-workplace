package itoa

import "testing"

func TestFormat(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{0, "0"}, {42, "42"}, {-17, "-17"}, {1000, "1000"},
	}
	for _, c := range cases {
		if got := Format(c.n); got != c.want {
			t.Errorf("Format(%d)=%q; want %q", c.n, got, c.want)
		}
	}
}
