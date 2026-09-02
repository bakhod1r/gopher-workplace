package lessergen

import "testing"

func TestMaxOf(t *testing.T) {
	got, ok := MaxOf([]Version{{1}, {3}, {2}})
	if !ok || got.N != 3 {
		t.Errorf("MaxOf = %+v, %v, want {3}, true", got, ok)
	}
}

func TestMaxOfTieKeepsEarlier(t *testing.T) {
	vs := []Version{{2}, {2}}
	got, _ := MaxOf(vs)
	if got != vs[0] {
		t.Errorf("MaxOf = %+v, want the first element", got)
	}
}

func TestMaxOfEmpty(t *testing.T) {
	got, ok := MaxOf([]Version{})
	if ok || got.N != 0 {
		t.Errorf("MaxOf(empty) = %+v, %v, want zero, false", got, ok)
	}
}

func TestMaxOfSingle(t *testing.T) {
	got, ok := MaxOf([]Version{{7}})
	if !ok || got.N != 7 {
		t.Errorf("MaxOf = %+v, %v, want {7}, true", got, ok)
	}
}
