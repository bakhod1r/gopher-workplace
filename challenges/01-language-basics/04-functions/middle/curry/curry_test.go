package curry

import "testing"

func TestAdd3(t *testing.T) {
	if got := Add3()(1)(2)(3); got != 6 {
		t.Errorf("=%d want 6", got)
	}
	if got := Add3()(10)(0)(0); got != 10 {
		t.Errorf("=%d want 10", got)
	}
}
