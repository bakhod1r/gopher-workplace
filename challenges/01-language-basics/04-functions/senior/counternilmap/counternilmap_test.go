package counternilmap

import "testing"

func TestTally(t *testing.T) {
	m := Tally([]string{"a", "b", "a", "a"})
	if m["a"] != 3 || m["b"] != 1 {
		t.Errorf("=%v want a:3 b:1", m)
	}
	if got := Tally(nil); len(got) != 0 {
		t.Errorf("empty input should give empty map")
	}
}
