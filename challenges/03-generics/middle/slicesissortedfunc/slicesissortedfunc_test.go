package slicesissortedfunc

import "testing"

func TestByName(t *testing.T) {
	if !ByName([]Item{{"a", 1}, {"b", 2}}) {
		t.Error("ByName(sorted) = false, want true")
	}
	if ByName([]Item{{"b", 1}, {"a", 2}}) {
		t.Error("ByName(unsorted) = true, want false")
	}
	if !ByName([]Item{{"a", 1}, {"a", 2}}) {
		t.Error("ByName(equal names) = false, want true")
	}
}

func TestByNameEdges(t *testing.T) {
	if !ByName(nil) {
		t.Error("ByName(nil) = false, want true")
	}
	if !ByName([]Item{{"only", 1}}) {
		t.Error("ByName(single) = false, want true")
	}
}
