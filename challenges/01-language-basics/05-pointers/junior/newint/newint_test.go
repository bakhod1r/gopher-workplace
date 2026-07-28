package newint

import "testing"

func TestAlloc(t *testing.T) {
	p := Alloc(7)
	if p == nil || *p != 7 {
		t.Errorf("*p wrong")
	}
	q := Alloc(7)
	if p == q {
		t.Errorf("each call should allocate a distinct int")
	}
}
