package clipcap

import "testing"

func TestClip(t *testing.T) {
	xs := make([]int, 3, 10)
	c := Clip(xs)
	if len(c) != 3 {
		t.Errorf("len=%d; want 3", len(c))
	}
	if cap(c) != 3 {
		t.Errorf("cap=%d; want 3 (clipped)", cap(c))
	}
}
