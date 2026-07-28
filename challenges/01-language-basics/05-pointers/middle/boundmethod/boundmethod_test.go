package boundmethod

import "testing"

func TestBind(t *testing.T) {
	c := &Counter{}
	inc := Bind(c)
	inc()
	inc()
	if c.N != 2 {
		t.Errorf("N=%d want 2", c.N)
	}
}
