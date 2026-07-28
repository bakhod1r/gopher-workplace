package methodexpr

import "testing"

func TestAdderExpr(t *testing.T) {
	f := AdderExpr()
	c := &Counter{}
	f(c, 5)
	f(c, 3)
	if c.N != 8 {
		t.Errorf("N=%d want 8 (receiver not applied?)", c.N)
	}
}
