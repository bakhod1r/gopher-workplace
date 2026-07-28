package mapbump

import "testing"

func TestBumpAll(t *testing.T) {
	a, b := 1, 2
	m := map[string]*int{"a": &a, "b": &b}
	BumpAll(m)
	if a != 2 || b != 3 {
		t.Errorf("a,b=%d,%d want 2,3", a, b)
	}
}
