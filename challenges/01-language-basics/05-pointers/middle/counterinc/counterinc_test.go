package counterinc

import "testing"

func TestInc(t *testing.T) {
	c := &Counter{}
	c.Inc()
	c.Inc()
	c.Inc()
	if c.N != 3 {
		t.Errorf("N=%d want 3", c.N)
	}
}
