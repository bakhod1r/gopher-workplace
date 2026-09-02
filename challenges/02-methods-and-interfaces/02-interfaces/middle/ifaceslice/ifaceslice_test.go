package ifaceslice

import "testing"

func TestID(t *testing.T) {
	if got := (User{ID_: "u1"}).ID(); got != "u1" {
		t.Errorf("User.ID = %q", got)
	}
	if got := (Order{ID_: "o1"}).ID(); got != "o1" {
		t.Errorf("Order.ID = %q", got)
	}
}

func TestToEntities(t *testing.T) {
	got := ToEntities([]User{{ID_: "u1"}, {ID_: "u2"}})
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].ID() != "u1" || got[1].ID() != "u2" {
		t.Errorf("ids = %v, %v", got[0].ID(), got[1].ID())
	}
	if n := len(ToEntities(nil)); n != 0 {
		t.Errorf("ToEntities(nil) len = %d, want 0", n)
	}
}

func TestIDsMixed(t *testing.T) {
	es := []Entity{User{ID_: "u1"}, Order{ID_: "o1"}}
	got := IDs(es)
	want := []string{"u1", "o1"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("IDs = %v, want %v", got, want)
		}
	}
	if n := len(IDs(nil)); n != 0 {
		t.Errorf("IDs(nil) len = %d, want 0", n)
	}
}

func TestEntitiesUsableTogether(t *testing.T) {
	es := ToEntities([]User{{ID_: "u1"}})
	es = append(es, Order{ID_: "o1"})
	if got := IDs(es); len(got) != 2 || got[1] != "o1" {
		t.Errorf("IDs = %v", got)
	}
}
