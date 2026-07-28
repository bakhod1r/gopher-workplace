package jsonescape

import "testing"

func TestEscape(t *testing.T) {
	cases := []struct{ in, want string }{
		{`plain`, `plain`},
		{`a"b`, `a\"b`},
		{`a\b`, `a\\b`}, // backslash must be escaped
		{"tab\there", `tab\there`},
		{"line\n", `line\n`},
	}
	for _, c := range cases {
		if got := Escape(c.in); got != c.want {
			t.Errorf("Escape(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
