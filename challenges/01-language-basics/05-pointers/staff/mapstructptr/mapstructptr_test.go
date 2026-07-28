package mapstructptr

import "testing"

func TestBumpVia(t *testing.T) {
	x := 10
	m := map[int]Ref{1: {P: &x}}
	BumpVia(m, 1)
	if x != 11 {
		t.Errorf("x=%d want 11 (didn't mutate through the pointer)", x)
	}
}
