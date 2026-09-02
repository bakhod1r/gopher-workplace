package ringgen

import (
	"reflect"
	"testing"
)

func TestRing(t *testing.T) {
	r := NewRing[int](2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Items() = %v, want []", got)
	}
	r.Add(1)
	r.Add(2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Items() = %v, want [1 2]", got)
	}
	r.Add(3)
	if got := r.Items(); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("Items() = %v, want [2 3] (oldest dropped)", got)
	}
}

func TestRingZeroSize(t *testing.T) {
	r := NewRing[int](0)
	r.Add(1)
	if got := r.Items(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Items() = %v, want []", got)
	}
}

func TestRingItemsIsACopy(t *testing.T) {
	r := NewRing[int](2)
	r.Add(1)
	got := r.Items()
	got[0] = 99
	if again := r.Items(); again[0] != 1 {
		t.Errorf("Items() returned the internal slice: %v, want [1]", again)
	}
}
