package rotate

import "testing"

func TestLeft(t *testing.T) {
	cases := []struct {
		b    byte
		n    int
		want byte
	}{
		{0b0000_0001, 1, 0b0000_0010},
		{0b1000_0000, 1, 0b0000_0001},
		{0b1010_0000, 4, 0b0000_1010},
		{0xFF, 3, 0xFF},
		{0b0000_0001, 8, 0b0000_0001},
		{0b0000_0001, 9, 0b0000_0010},
	}
	for _, c := range cases {
		if got := Left(c.b, c.n); got != c.want {
			t.Errorf("Left(%#08b,%d)=%#08b; want %#08b", c.b, c.n, got, c.want)
		}
	}
}
