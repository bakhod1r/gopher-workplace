package sortbykeygen

import "testing"

type person struct {
	name string
	age  int
}

func ageOf(p person) int { return p.age }

func TestSortedBy(t *testing.T) {
	in := []person{{"a", 30}, {"b", 20}, {"c", 25}}
	got := SortedBy(in, ageOf)
	want := []string{"b", "c", "a"}
	for i, w := range want {
		if got[i].name != w {
			t.Fatalf("SortedBy = %+v, want order %v", got, want)
		}
	}
	if in[0].name != "a" {
		t.Errorf("input mutated: %+v", in)
	}
}

func TestSortedByIsStable(t *testing.T) {
	in := []person{{"a", 20}, {"b", 20}, {"c", 10}}
	got := SortedBy(in, ageOf)
	if got[1].name != "a" || got[2].name != "b" {
		t.Errorf("equal keys reordered: %+v", got)
	}
}

func TestSortedByEmpty(t *testing.T) {
	got := SortedBy([]person(nil), ageOf)
	if got == nil || len(got) != 0 {
		t.Errorf("SortedBy(nil) = %v, want an empty non-nil slice", got)
	}
}
