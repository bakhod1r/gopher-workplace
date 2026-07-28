package area

import "testing"

func TestArea(t *testing.T) {
	cases := []struct {
		w, h int32
		want int64
	}{
		{3, 4, 12},
		{100000, 100000, 10000000000}, // overflows int32
		{46341, 46341, 2147488281},    // just over 2^31
	}
	for _, c := range cases {
		if got := Area(c.w, c.h); got != c.want {
			t.Errorf("Area(%d,%d)=%d; want %d", c.w, c.h, got, c.want)
		}
	}
}
