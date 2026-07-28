package sign

import "testing"

func TestNegative(t *testing.T) {
	cases := []struct {
		b   uint8
		neg bool
	}{
		{0x00, false},
		{0x7F, false}, // 127
		{0x80, true},  // -128
		{0xFF, true},  // -1
		{0x40, false}, // 64
	}
	for _, c := range cases {
		if got := Negative(c.b); got != c.neg {
			t.Errorf("Negative(%#x)=%v; want %v", c.b, got, c.neg)
		}
	}
}
