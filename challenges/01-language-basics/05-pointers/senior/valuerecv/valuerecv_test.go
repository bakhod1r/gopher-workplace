package valuerecv

import "testing"

func TestInc(t *testing.T) {
	c := &Counter{}
	c.Inc()
	c.Inc()
	if c.N != 2 {
		t.Errorf("N=%d want 2 (value receiver lost the increment)", c.N)
	}
}
