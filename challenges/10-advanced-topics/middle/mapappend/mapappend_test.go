package mapappend

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := map[string][]int{}
	Add(m, "a", 1)
	Add(m, "a", 2)
	Add(m, "b", 3)
	want := map[string][]int{"a": {1, 2}, "b": {3}}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("m = %v, want %v", m, want)
	}
}

func TestAddCreatesTheEntry(t *testing.T) {
	m := map[string][]int{}
	Add(m, "new", 7)
	got, ok := m["new"]
	if !ok {
		t.Fatal("the key was not created")
	}
	if !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("m[new] = %v, want [7]", got)
	}
}

func TestAddNilMap(t *testing.T) {
	Add(nil, "a", 1)
}

func TestAddManyValues(t *testing.T) {
	m := map[string][]int{}
	for i := 0; i < 100; i++ {
		Add(m, "k", i)
	}
	if len(m["k"]) != 100 {
		t.Errorf("len = %d, want 100", len(m["k"]))
	}
	for i, v := range m["k"] {
		if v != i {
			t.Fatalf("m[k][%d] = %d, want %d", i, v, i)
		}
	}
}
