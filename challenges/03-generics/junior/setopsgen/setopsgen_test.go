package setopsgen

import "testing"

func set(vs ...int) map[int]struct{} {
	m := make(map[int]struct{})
	for _, v := range vs {
		m[v] = struct{}{}
	}
	return m
}

func keys(m map[int]struct{}) []int {
	out := make([]int, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func has(m map[int]struct{}, v int) bool {
	_, ok := m[v]
	return ok
}

func TestUnion(t *testing.T) {
	u := Union(set(1, 2), set(2, 3))
	if len(u) != 3 || !has(u, 1) || !has(u, 2) || !has(u, 3) {
		t.Errorf("Union = %v, want {1 2 3}", keys(u))
	}
}

func TestIntersect(t *testing.T) {
	i := Intersect(set(1, 2), set(2, 3))
	if len(i) != 1 || !has(i, 2) {
		t.Errorf("Intersect = %v, want {2}", keys(i))
	}
	e := Intersect(set(1), set())
	if e == nil || len(e) != 0 {
		t.Errorf("Intersect with an empty set = %v, want an empty non-nil map", keys(e))
	}
}

func TestSetOpsDoNotMutate(t *testing.T) {
	a := set(1)
	b := set(2)
	Union(a, b)
	Intersect(a, b)
	if len(a) != 1 || len(b) != 1 {
		t.Errorf("inputs mutated: a=%v b=%v", keys(a), keys(b))
	}
}
