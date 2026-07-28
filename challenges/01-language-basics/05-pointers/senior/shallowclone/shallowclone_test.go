package shallowclone

import "testing"

func TestClone(t *testing.T) {
	d := &Doc{Tags: []string{"a", "b"}}
	c := Clone(d)
	c.Tags[0] = "X"
	if d.Tags[0] != "a" {
		t.Errorf("original mutated: %v (shared slice)", d.Tags)
	}
}
