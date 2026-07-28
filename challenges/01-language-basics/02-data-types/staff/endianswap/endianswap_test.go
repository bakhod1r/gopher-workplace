package endianswap

import "testing"

func TestSwap32(t *testing.T) {
	cases := []struct {
		x, want uint32
	}{
		{0x11223344, 0x44332211},
		{0x00000001, 0x01000000},
		{0xDEADBEEF, 0xEFBEADDE},
		{0, 0},
	}
	for _, c := range cases {
		if got := Swap32(c.x); got != c.want {
			t.Errorf("Swap32(%#08x)=%#08x; want %#08x", c.x, got, c.want)
		}
	}
}
