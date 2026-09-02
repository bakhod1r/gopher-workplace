package multimapgen

import (
	"reflect"
	"testing"
)

func TestMultiMap(t *testing.T) {
	m := NewMultiMap[string, int]()
	if got := m.Get("missing"); !reflect.DeepEqual(got, []int{}) {
		t.Errorf(`Get("missing") = %v, want []`, got)
	}
	m.Add("a", 1)
	m.Add("a", 2)
	m.Add("b", 3)
	if got := m.Get("a"); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf(`Get("a") = %v, want [1 2]`, got)
	}
	if got := m.Get("b"); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf(`Get("b") = %v, want [3]`, got)
	}
}

func TestMultiMapGetIsACopy(t *testing.T) {
	m := NewMultiMap[string, int]()
	m.Add("a", 1)
	got := m.Get("a")
	got[0] = 99
	if again := m.Get("a"); again[0] != 1 {
		t.Errorf("Get returned internal storage: %v, want [1]", again)
	}
}
