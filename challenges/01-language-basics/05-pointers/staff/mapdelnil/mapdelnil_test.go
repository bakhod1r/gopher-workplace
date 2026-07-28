package mapdelnil

import "testing"

func TestRemove(t *testing.T) {
	x := 5
	m := map[int]*int{1: &x, 2: &x}
	Remove(m, 1)
	if len(m) != 1 {
		t.Errorf("len=%d want 1 (key still present)", len(m))
	}
	if _, ok := m[1]; ok {
		t.Errorf("key 1 should be gone")
	}
}
