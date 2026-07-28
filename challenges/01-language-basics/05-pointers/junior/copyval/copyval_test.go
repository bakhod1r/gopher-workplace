package copyval

import "testing"

func TestCopyInto(t *testing.T) {
	a, b := 1, 9
	CopyInto(&a, &b)
	if a != 9 {
		t.Errorf("a=%d want 9", a)
	}
	if b != 9 {
		t.Errorf("b changed: %d", b)
	}
}
