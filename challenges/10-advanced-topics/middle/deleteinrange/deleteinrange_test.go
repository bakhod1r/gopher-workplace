package deleteinrange

import (
	"reflect"
	"testing"
)

func TestRemoveEven(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	if got := RemoveEven(m); got != 2 {
		t.Errorf("RemoveEven = %d, want 2", got)
	}
	want := map[int]int{1: 1, 3: 3}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("m = %v, want %v", m, want)
	}
}

func TestRemoveEvenNegativeKeys(t *testing.T) {
	m := map[int]int{-2: 1, -1: 1, 0: 1}
	if got := RemoveEven(m); got != 2 {
		t.Errorf("RemoveEven = %d, want 2: -2 and 0 are even", got)
	}
	if _, ok := m[-1]; !ok || len(m) != 1 {
		t.Errorf("m = %v, want map[-1:1]", m)
	}
}

func TestRemoveEvenEmpty(t *testing.T) {
	if got := RemoveEven(nil); got != 0 {
		t.Errorf("RemoveEven(nil) = %d, want 0", got)
	}
	m := map[int]int{}
	if got := RemoveEven(m); got != 0 {
		t.Errorf("RemoveEven = %d, want 0", got)
	}
}

func TestRemoveEvenIsVisibleToTheCaller(t *testing.T) {
	m := map[int]int{2: 2}
	alias := m
	RemoveEven(m)
	if len(alias) != 0 {
		t.Error("the deletion was not applied to the caller's map")
	}
}

func TestRemoveEvenLarge(t *testing.T) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	if got := RemoveEven(m); got != 500 {
		t.Errorf("RemoveEven = %d, want 500", got)
	}
	if len(m) != 500 {
		t.Errorf("len = %d, want 500", len(m))
	}
	for k := range m {
		if k%2 == 0 {
			t.Fatalf("key %d survived", k)
		}
	}
}
