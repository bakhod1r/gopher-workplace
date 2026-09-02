package maxbykey

import "testing"

type person struct {
	name  string
	score int
}

func score(p person) int { return p.score }

func TestMaxBy(t *testing.T) {
	people := []person{{"a", 1}, {"b", 3}, {"c", 2}}
	got, ok := MaxBy(people, score)
	if !ok || got.name != "b" {
		t.Errorf("MaxBy = %+v, %v, want {b 3}, true", got, ok)
	}

	tie := []person{{"a", 2}, {"b", 2}}
	got, ok = MaxBy(tie, score)
	if !ok || got.name != "a" {
		t.Errorf("MaxBy(tie) = %+v, %v, want {a 2}, true (earlier wins)", got, ok)
	}

	got, ok = MaxBy([]person{}, score)
	if ok || got.name != "" {
		t.Errorf("MaxBy(empty) = %+v, %v, want zero, false", got, ok)
	}
}

func TestMaxByStringKey(t *testing.T) {
	got, ok := MaxBy([]person{{"a", 1}, {"z", 0}}, func(p person) string { return p.name })
	if !ok || got.name != "z" {
		t.Errorf("MaxBy by name = %+v, %v, want {z 0}, true", got, ok)
	}
}
