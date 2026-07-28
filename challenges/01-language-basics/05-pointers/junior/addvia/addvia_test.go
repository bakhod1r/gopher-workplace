package addvia

import "testing"

func TestAdd(t *testing.T) {
	x := 10
	Add(&x, 5)
	Add(&x, -3)
	if x != 12 {
		t.Errorf("x=%d want 12", x)
	}
}
