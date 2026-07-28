package mapclonesharedslice

import "testing"

func TestClone(t *testing.T) {
	m := map[string][]int{"a": {1, 2}, "b": {3}}
	c := Clone(m)
	c["a"][0] = 99
	if m["a"][0] != 1 {
		t.Errorf("slice value shared: m[a]=%v", m["a"])
	}
	if len(c) != 2 || c["b"][0] != 3 {
		t.Errorf("clone wrong: %v", c)
	}
}
