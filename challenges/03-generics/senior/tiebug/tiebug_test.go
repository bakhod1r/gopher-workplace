package tiebug

import "testing"

type row struct {
	name  string
	score int
}

func score(r row) int { return r.score }

func TestMaxByTie(t *testing.T) {
	rows := []row{{"a", 3}, {"b", 3}}
	got, ok := MaxBy(rows, score)
	if !ok || got.name != "a" {
		t.Errorf("MaxBy = %+v, %v, want {a 3}, true (earlier wins ties)", got, ok)
	}
}

func TestMaxByPicksMaximum(t *testing.T) {
	rows := []row{{"a", 1}, {"b", 3}, {"c", 2}}
	got, _ := MaxBy(rows, score)
	if got.name != "b" {
		t.Errorf("MaxBy = %+v, want {b 3}", got)
	}
}

func TestMaxByEmpty(t *testing.T) {
	got, ok := MaxBy([]row{}, score)
	if ok || got.name != "" {
		t.Errorf("MaxBy(empty) = %+v, %v, want zero, false", got, ok)
	}
}
