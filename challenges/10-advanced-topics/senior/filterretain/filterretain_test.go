package filterretain

import "testing"

func TestKeepSelects(t *testing.T) {
	in := []Record{{ID: 1, Size: 10}, {ID: 2, Size: 200}, {ID: 3, Size: 300}}
	got := Keep(in, 100)
	if len(got) != 2 || got[0].ID != 2 || got[1].ID != 3 {
		t.Fatalf("Keep = %v, want records 2 and 3", got)
	}
	if got := Keep(nil, 1); len(got) != 0 {
		t.Errorf("Keep(nil) = %v, want empty", got)
	}
	if got := Keep([]Record{{Size: 1}}, 100); len(got) != 0 {
		t.Errorf("Keep = %v, want empty", got)
	}
}

func TestKeepDoesNotMutateTheInput(t *testing.T) {
	in := []Record{{ID: 1, Size: 10}, {ID: 2, Size: 200}}
	Keep(in, 100)
	if in[0].ID != 1 {
		t.Errorf("in[0].ID = %d, want 1: the batch was rewritten in place", in[0].ID)
	}
}

func TestKeepReleasesTheBatch(t *testing.T) {
	in := make([]Record, 1<<14)
	for i := range in {
		in[i] = Record{ID: i, Size: i}
	}
	got := Keep(in, 1<<14-4)
	if cap(got) > 64 {
		t.Errorf("cap = %d, want a right-sized result: it still owns the batch's array", cap(got))
	}
}
