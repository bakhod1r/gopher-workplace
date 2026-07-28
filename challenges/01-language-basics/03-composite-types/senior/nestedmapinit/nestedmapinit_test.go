package nestedmapinit

import "testing"

func TestTally(t *testing.T) {
	m := Tally([][2]string{{"a", "x"}, {"a", "x"}, {"a", "y"}, {"b", "x"}})
	if m["a"]["x"] != 2 || m["a"]["y"] != 1 || m["b"]["x"] != 1 {
		t.Errorf("got %v", m)
	}
}
