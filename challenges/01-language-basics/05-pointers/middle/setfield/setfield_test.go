package setfield

import "testing"

func TestSetName(t *testing.T) {
	u := &User{Name: "old"}
	m := map[int]*User{1: u}
	if !SetName(m, 1, "new") {
		t.Errorf("should succeed")
	}
	if u.Name != "new" {
		t.Errorf("Name=%q want new", u.Name)
	}
	if SetName(m, 2, "x") {
		t.Errorf("absent id should return false")
	}
}
