package stablebug

import "testing"

type person struct {
	name string
	age  int
}

func ageOf(p person) int { return p.age }

func TestSortedByIsStable(t *testing.T) {
	in := []person{
		{"a", 20}, {"b", 20}, {"c", 10}, {"d", 20}, {"e", 10},
		{"f", 20}, {"g", 10}, {"h", 20}, {"i", 10}, {"j", 20},
		{"k", 20}, {"l", 10}, {"m", 20}, {"n", 10}, {"o", 20},
	}
	got := SortedBy(in, ageOf)

	var tens, twenties []string
	for _, p := range got {
		if p.age == 10 {
			tens = append(tens, p.name)
		} else {
			twenties = append(twenties, p.name)
		}
	}
	if len(got) != len(in) {
		t.Fatalf("len = %d, want %d", len(got), len(in))
	}
	if got[0].age != 10 || got[len(got)-1].age != 20 {
		t.Fatalf("not sorted by age: %+v", got)
	}

	wantTens := []string{"c", "e", "g", "i", "l", "n"}
	wantTwenties := []string{"a", "b", "d", "f", "h", "j", "k", "m", "o"}
	for i, w := range wantTens {
		if tens[i] != w {
			t.Fatalf("age 10 group = %v, want %v (equal keys must keep input order)", tens, wantTens)
		}
	}
	for i, w := range wantTwenties {
		if twenties[i] != w {
			t.Fatalf("age 20 group = %v, want %v", twenties, wantTwenties)
		}
	}
}

func TestSortedByDoesNotMutate(t *testing.T) {
	in := []person{{"a", 30}, {"b", 10}}
	SortedBy(in, ageOf)
	if in[0].name != "a" {
		t.Errorf("input mutated: %+v", in)
	}
	got := SortedBy([]person(nil), ageOf)
	if got == nil || len(got) != 0 {
		t.Errorf("SortedBy(nil) = %v, want an empty non-nil slice", got)
	}
}
