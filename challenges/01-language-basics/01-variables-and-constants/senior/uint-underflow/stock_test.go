package stock

import "testing"

func TestRemaining(t *testing.T) {
	cases := []struct {
		have, sold, want uint
	}{
		{10, 3, 7},
		{5, 5, 0},
		{2, 9, 0}, // oversell must clamp, not wrap to a huge uint
	}
	for _, c := range cases {
		if got := Remaining(c.have, c.sold); got != c.want {
			t.Errorf("Remaining(%d,%d)=%d; want %d", c.have, c.sold, got, c.want)
		}
	}
}
