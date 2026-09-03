package stdsortstableweakbug

import (
	"testing"
	"time"
)

func names(ts []Task) []string {
	out := make([]string, len(ts))
	for i, t := range ts {
		out[i] = t.Name
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

func TestSortByPriorityOrders(t *testing.T) {
	got := SortByPriority([]Task{{"a", 2}, {"b", 1}, {"c", 2}})
	if !eq(names(got), "b", "a", "c") {
		t.Errorf("SortByPriority = %v, want [b a c]", names(got))
	}
}

func TestSortByPriorityIsStable(t *testing.T) {
	got := SortByPriority([]Task{{"a", 1}, {"b", 1}, {"c", 1}})
	if !eq(names(got), "a", "b", "c") {
		t.Errorf("SortByPriority = %v, want [a b c]", names(got))
	}
}

func TestSortByPriorityLeavesInput(t *testing.T) {
	in := []Task{{"a", 2}, {"b", 1}}
	SortByPriority(in)
	if !eq(names(in), "a", "b") {
		t.Errorf("input reordered: %v", names(in))
	}
}

func TestSortByPriorityAtScale(t *testing.T) {
	const n = 200000
	in := make([]Task, n)
	for i := range in {
		in[i] = Task{Pri: i % 50}
	}
	start := time.Now()
	got := SortByPriority(in)
	if len(got) != n {
		t.Fatalf("length = %d, want %d", len(got), n)
	}
	for i := 1; i < len(got); i++ {
		if got[i-1].Pri > got[i].Pri {
			t.Fatalf("not sorted at %d: %d before %d", i, got[i-1].Pri, got[i].Pri)
		}
	}
	if d := time.Since(start); d > 5*time.Second {
		t.Fatalf("scale sort took %v, want under 5s", d)
	}
}
