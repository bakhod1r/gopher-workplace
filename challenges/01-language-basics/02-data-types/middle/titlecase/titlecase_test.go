package titlecase

import "testing"

func TestTitle(t *testing.T) {
	cases := []struct{ in, want string }{
		{"hello world", "Hello World"},
		{"GO is FUN", "Go Is Fun"},
		{"a", "A"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Title(c.in); got != c.want {
			t.Errorf("Title(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
