package normspaces

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct{ in, want string }{
		{"  hello   world  ", "hello world"},
		{"a\tb\nc", "a b c"},
		{"single", "single"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Normalize(c.in); got != c.want {
			t.Errorf("Normalize(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
