package uniquekeybug

import "testing"

type rec struct {
	id  int
	tag string
}

func idOf(r rec) int { return r.id }

func TestUniqueByKeepsFirst(t *testing.T) {
	rows := []rec{{1, "a"}, {1, "b"}, {2, "c"}}
	got := UniqueBy(rows, idOf)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].tag != "a" {
		t.Errorf("kept %q for id 1, want a (the first occurrence)", got[0].tag)
	}
	if got[1].tag != "c" {
		t.Errorf("second element = %+v, want {2 c}", got[1])
	}
}

func TestUniqueByOrder(t *testing.T) {
	rows := []rec{{2, "x"}, {1, "y"}, {2, "z"}}
	got := UniqueBy(rows, idOf)
	if len(got) != 2 || got[0].id != 2 || got[1].id != 1 {
		t.Errorf("UniqueBy = %+v, want ids [2 1]", got)
	}
	if got[0].tag != "x" {
		t.Errorf("kept %q for id 2, want x", got[0].tag)
	}
}

func TestUniqueByEmpty(t *testing.T) {
	got := UniqueBy([]rec{}, idOf)
	if got == nil || len(got) != 0 {
		t.Errorf("UniqueBy(empty) = %v, want an empty non-nil slice", got)
	}
}
