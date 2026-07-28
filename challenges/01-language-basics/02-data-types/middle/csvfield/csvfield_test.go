package csvfield

import "testing"

func TestQuote(t *testing.T) {
	cases := []struct{ in, want string }{
		{"plain", "plain"},
		{"a,b", "\"a,b\""},
		{"say \"hi\"", "\"say \"\"hi\"\"\""},
		{"line1\nline2", "\"line1\nline2\""},
		{"", ""},
	}
	for _, c := range cases {
		if got := Quote(c.in); got != c.want {
			t.Errorf("Quote(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
