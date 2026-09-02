package unionfindgen

import "testing"

func TestUnionConnected(t *testing.T) {
	d := NewDisjoint[string]()
	if d.Connected("a", "b") {
		t.Error(`Connected("a", "b") = true, want false`)
	}
	d.Union("a", "b")
	if !d.Connected("a", "b") {
		t.Error(`Connected("a", "b") = false, want true`)
	}
	if d.Connected("a", "c") {
		t.Error(`Connected("a", "c") = true, want false`)
	}
}

func TestUnionTransitive(t *testing.T) {
	d := NewDisjoint[int]()
	d.Union(1, 2)
	d.Union(2, 3)
	if !d.Connected(1, 3) {
		t.Error("Connected(1, 3) = false, want true")
	}
}

func TestFindIsStable(t *testing.T) {
	d := NewDisjoint[int]()
	if d.Find(7) != 7 {
		t.Error("Find on an unseen element should return the element itself")
	}
	d.Union(1, 2)
	if d.Find(1) != d.Find(2) {
		t.Error("Find(1) and Find(2) differ after Union")
	}
	if d.Find(1) != d.Find(1) {
		t.Error("Find is not stable across calls")
	}
}
