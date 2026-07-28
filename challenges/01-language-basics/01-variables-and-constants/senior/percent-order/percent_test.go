package percent

import "testing"

func TestPercent(t *testing.T) {
	cases := []struct {
		part, total, want int
	}{
		{1, 4, 25},
		{3, 4, 75},
		{1, 3, 33},
		{50, 50, 100},
	}
	for _, c := range cases {
		if got := Percent(c.part, c.total); got != c.want {
			t.Errorf("Percent(%d,%d)=%d; want %d", c.part, c.total, got, c.want)
		}
	}
}
