package nestptr

import "testing"

func TestAdd(t *testing.T) {
	m := map[string]map[string]int{}
	Add(m, "g", "a")
	Add(m, "g", "a")
	Add(m, "g", "b")
	if m["g"]["a"] != 2 || m["g"]["b"] != 1 {
		t.Errorf("counts wrong: %v", m["g"])
	}
}
