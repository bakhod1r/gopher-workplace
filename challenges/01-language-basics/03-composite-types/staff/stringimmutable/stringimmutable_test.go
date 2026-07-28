package stringimmutable

import "testing"

func TestUpper(t *testing.T) {
	cases := []struct{ in, want string }{
		{"hello", "HELLO"},
		{"Go123", "GO123"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Upper(c.in); got != c.want {
			t.Errorf("Upper(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
