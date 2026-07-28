package unsafeslice

import "testing"

func TestView(t *testing.T) {
	a := [4]int32{1, 2, 3, 4}
	v := View(&a)
	if len(v) != 4 {
		t.Fatalf("len=%d want 4 (byte length instead of element count?)", len(v))
	}
	if v[3] != 4 {
		t.Fatalf("v[3]=%d want 4", v[3])
	}
}
