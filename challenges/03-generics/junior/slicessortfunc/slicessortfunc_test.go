package slicessortfunc

import "testing"

func TestSortByAge(t *testing.T) {
	people := []Person{{"a", 30}, {"b", 20}, {"c", 25}}
	SortByAge(people)
	want := []string{"b", "c", "a"}
	for i, w := range want {
		if people[i].Name != w {
			t.Errorf("people[%d].Name = %q, want %q (order: %+v)", i, people[i].Name, w, people)
			break
		}
	}
}

func TestSortByAgeIsStable(t *testing.T) {
	people := []Person{{"a", 20}, {"b", 20}, {"c", 10}}
	SortByAge(people)
	if people[1].Name != "a" || people[2].Name != "b" {
		t.Errorf("equal ages reordered: %+v, want c, a, b", people)
	}
}

func TestSortByAgeEmpty(t *testing.T) {
	var people []Person
	SortByAge(people)
	if len(people) != 0 {
		t.Errorf("empty slice changed: %+v", people)
	}
}
