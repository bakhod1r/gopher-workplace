package incptr

import "testing"

func TestInc(t *testing.T) {
	x := 41
	Inc(&x)
	if x != 42 {
		t.Errorf("x=%d want 42", x)
	}
}
