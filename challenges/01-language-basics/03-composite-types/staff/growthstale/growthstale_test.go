package growthstale

import "testing"

func TestBumpFirst(t *testing.T) {
	first, s := BumpFirst(20)
	if first != 99 {
		t.Errorf("s[0]=%d; want 99 (stale pointer wrote the old array)", first)
	}
	if len(s) != 2 || s[1] != 20 {
		t.Errorf("s=%v; want [99 20]", s)
	}
}
