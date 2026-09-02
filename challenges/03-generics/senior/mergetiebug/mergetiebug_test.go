package mergetiebug

import "testing"

type rec struct {
	src string
	k   int
}

func key(r rec) int { return r.k }

func srcs(rs []rec) []string {
	out := make([]string, len(rs))
	for i, r := range rs {
		out[i] = r.src
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

func TestMergeSortedOrder(t *testing.T) {
	a := []rec{{"a1", 1}, {"a3", 3}}
	b := []rec{{"b2", 2}}
	got := MergeSorted(a, b, key)
	if !eq(srcs(got), "a1", "b2", "a3") {
		t.Errorf("MergeSorted = %v, want [a1 b2 a3]", srcs(got))
	}
	if got := MergeSorted(nil, b, key); len(got) != 1 {
		t.Errorf("MergeSorted = %v, want one element", got)
	}
}

func TestMergeSortedTieTakesFromA(t *testing.T) {
	a := []rec{{"fromA", 2}}
	b := []rec{{"fromB", 2}}
	got := MergeSorted(a, b, key)
	if !eq(srcs(got), "fromA", "fromB") {
		t.Errorf("MergeSorted = %v, want [fromA fromB] (a wins the tie)", srcs(got))
	}
}
