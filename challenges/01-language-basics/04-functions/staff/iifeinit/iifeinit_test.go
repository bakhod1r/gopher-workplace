package iifeinit

import "testing"

func TestBuildTable(t *testing.T) {
	m := BuildTable(4)
	for i := 0; i < 4; i++ {
		if m[i] != i*i {
			t.Errorf("m[%d]=%d want %d", i, m[i], i*i)
		}
	}
	if len(m) != 4 {
		t.Errorf("len=%d want 4", len(m))
	}
}
