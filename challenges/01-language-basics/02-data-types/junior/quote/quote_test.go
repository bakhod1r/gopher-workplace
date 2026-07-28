package quote

import "testing"

func TestWrap(t *testing.T) {
	cases := []struct{ in, want string }{
		{"hello", "\"hello\""},
		{"", "\"\""},
		{"a b", "\"a b\""},
	}
	for _, c := range cases {
		if got := Wrap(c.in); got != c.want {
			t.Errorf("Wrap(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
