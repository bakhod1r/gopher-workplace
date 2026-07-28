package hexdigit

import "testing"

func TestDigit(t *testing.T) {
	cases := []struct {
		n    int
		want byte
	}{
		{0, '0'}, {9, '9'}, {10, 'a'}, {15, 'f'}, {16, '?'}, {-1, '?'},
	}
	for _, c := range cases {
		if got := Digit(c.n); got != c.want {
			t.Errorf("Digit(%d)=%q; want %q", c.n, got, c.want)
		}
	}
}
