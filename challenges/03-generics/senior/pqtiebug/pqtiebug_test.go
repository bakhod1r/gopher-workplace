package pqtiebug

import "testing"

type job struct {
	name string
	prio int
}

func prio(j job) int { return j.prio }

func names(js []job) []string {
	out := make([]string, len(js))
	for i, j := range js {
		out[i] = j.name
	}
	return out
}

func eq(a []string, b ...string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInsertSortedPlacement(t *testing.T) {
	s := []job{{"a", 1}, {"c", 3}}
	got := InsertSorted(s, job{"b", 2}, prio)
	if !eq(names(got), "a", "b", "c") {
		t.Errorf("InsertSorted = %v, want [a b c]", names(got))
	}
	if got := InsertSorted(nil, job{"x", 5}, prio); len(got) != 1 {
		t.Errorf("InsertSorted = %v, want one element", got)
	}
}

func TestInsertSortedGoesAfterEquals(t *testing.T) {
	s := []job{{"first", 2}, {"later", 3}}
	got := InsertSorted(s, job{"second", 2}, prio)
	if !eq(names(got), "first", "second", "later") {
		t.Errorf("InsertSorted = %v, want [first second later]", names(got))
	}
}
