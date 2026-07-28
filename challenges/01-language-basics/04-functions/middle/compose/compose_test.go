package compose

import "testing"

func TestCompose(t *testing.T) {
	inc := func(x int) int { return x + 1 }
	dbl := func(x int) int { return x * 2 }
	h := Compose(inc, dbl) // inc(dbl(x))
	if h(3) != 7 {
		t.Errorf("Compose(inc,dbl)(3)=%d want 7", h(3))
	}
	h2 := Compose(dbl, inc) // dbl(inc(x))
	if h2(3) != 8 {
		t.Errorf("=%d want 8", h2(3))
	}
}
