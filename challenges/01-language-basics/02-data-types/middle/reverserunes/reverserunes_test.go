package reverserunes

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct{ in, want string }{
		{"hello", "olleh"},
		{"café", "éfac"},
		{"日本語", "語本日"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
