package treemapgen

import (
	"reflect"
	"testing"
)

func TestSortedMapKeys(t *testing.T) {
	m := NewSorted[string, int]()
	m.Set("b", 1)
	m.Set("a", 2)
	m.Set("c", 3)
	if got := m.Keys(); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Keys() = %v, want [a b c]", got)
	}
}

func TestSortedMapUpdate(t *testing.T) {
	m := NewSorted[int, string]()
	m.Set(2, "x")
	m.Set(1, "y")
	m.Set(2, "z")
	if got := m.Keys(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Keys() = %v, want [1 2]", got)
	}
	if v, ok := m.Get(2); v != "z" || !ok {
		t.Errorf("Get(2) = %q, %v, want z, true", v, ok)
	}
}

func TestSortedMapMissing(t *testing.T) {
	m := NewSorted[string, int]()
	if v, ok := m.Get("nope"); v != 0 || ok {
		t.Errorf("Get(miss) = %v, %v, want 0, false", v, ok)
	}
	if got := m.Keys(); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("Keys() = %v, want []", got)
	}
}
