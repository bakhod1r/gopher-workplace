package detmapgen

import "testing"

func TestEntriesSorted(t *testing.T) {
	got := Entries(map[string]int{"b": 1, "a": 2, "c": 3})
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	want := []string{"a", "b", "c"}
	for i, w := range want {
		if got[i].Key != w {
			t.Fatalf("Entries = %+v, want keys %v", got, want)
		}
	}
	if got[0].Value != 2 {
		t.Errorf("value for a = %v, want 2", got[0].Value)
	}
}

func TestEntriesIsDeterministic(t *testing.T) {
	m := map[int]string{3: "c", 1: "a", 2: "b"}
	first := Entries(m)
	for i := 0; i < 30; i++ {
		again := Entries(m)
		for j := range first {
			if again[j].Key != first[j].Key {
				t.Fatalf("run %d differs: %+v vs %+v", i, again, first)
			}
		}
	}
}

func TestEntriesEmpty(t *testing.T) {
	got := Entries(map[string]int(nil))
	if got == nil || len(got) != 0 {
		t.Errorf("Entries(nil) = %v, want an empty non-nil slice", got)
	}
}
