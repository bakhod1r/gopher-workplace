package clampbyte

import "testing"

func TestClamp(t *testing.T) {
	cases := []struct {
		x    int
		want byte
	}{
		{128, 128}, {0, 0}, {255, 255}, {300, 255}, {-20, 0},
	}
	for _, c := range cases {
		if got := Clamp(c.x); got != c.want {
			t.Errorf("Clamp(%d)=%d; want %d", c.x, got, c.want)
		}
	}
}
