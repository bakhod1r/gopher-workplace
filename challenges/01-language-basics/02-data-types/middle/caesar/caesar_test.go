package caesar

import "testing"

func TestShift(t *testing.T) {
	cases := []struct {
		s    string
		n    int
		want string
	}{
		{"abc", 1, "bcd"},
		{"xyz", 3, "abc"},
		{"Hello, World!", 13, "Uryyb, Jbeyq!"},
		{"abc", -1, "zab"},
	}
	for _, c := range cases {
		if got := Shift(c.s, c.n); got != c.want {
			t.Errorf("Shift(%q,%d)=%q; want %q", c.s, c.n, got, c.want)
		}
	}
}
