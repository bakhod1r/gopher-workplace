package runlength

import "testing"

func TestEncode(t *testing.T) {
	cases := []struct{ in, want string }{
		{"aaab", "a3b1"},
		{"abc", "a1b1c1"},
		{"aaaa", "a4"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Encode(c.in); got != c.want {
			t.Errorf("Encode(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
