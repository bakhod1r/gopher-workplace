package nestedmap

import "testing"

func TestAdd(t *testing.T) {
	g := map[string]map[string]bool{}
	Add(g, "a", "b")
	Add(g, "a", "c")
	if !g["a"]["b"] || !g["a"]["c"] {
		t.Errorf("edges not recorded: %v", g)
	}
}
