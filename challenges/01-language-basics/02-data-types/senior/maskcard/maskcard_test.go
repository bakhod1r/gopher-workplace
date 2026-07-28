package maskcard

import "testing"

func TestMask(t *testing.T) {
	cases := []struct{ in, want string }{
		{"1234567890123456", "************3456"},
		{"4242", "4242"},
		{"12345", "*2345"},
	}
	for _, c := range cases {
		if got := Mask(c.in); got != c.want {
			t.Errorf("Mask(%q)=%q; want %q", c.in, got, c.want)
		}
	}
}
