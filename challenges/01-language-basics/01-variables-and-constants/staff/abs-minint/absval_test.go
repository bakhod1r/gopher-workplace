package absval

import "testing"

func TestAbs(t *testing.T) {
	cases := []struct {
		x    int8
		want int
	}{
		{5, 5}, {-5, 5}, {0, 0}, {-128, 128}, {127, 127},
	}
	for _, c := range cases {
		if got := Abs(c.x); got != c.want {
			t.Errorf("Abs(%d)=%d; want %d", c.x, got, c.want)
		}
	}
}
