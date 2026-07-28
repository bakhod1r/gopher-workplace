package tolowerascii

import "testing"

func TestLower(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Hello", "hello"},
		{"ABC123", "abc123"},
		{"a[b]", "a[b]"}, // '[' (91) is just past 'Z' (90) — must stay
		{"MixED_case", "mixed_case"},
	}
	for _, c := range cases {
		if got := Lower(c.in); got != c.want {
			t.Errorf("Lower(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
