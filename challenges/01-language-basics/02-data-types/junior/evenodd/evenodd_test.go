package evenodd

import "testing"

func TestParity(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{0, "even"}, {2, "even"}, {3, "odd"}, {-4, "even"}, {-7, "odd"},
	}
	for _, c := range cases {
		if got := Parity(c.n); got != c.want {
			t.Errorf("Parity(%d)=%q; want %q", c.n, got, c.want)
		}
	}
}
