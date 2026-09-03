package ringstartbug

import (
	"reflect"
	"testing"
)

func TestRingPartialAfterWrap(t *testing.T) {
	r := NewRing[int](3)
	for i := 1; i <= 4; i++ {
		r.Add(i)
	}
	if got := r.Items(); !reflect.DeepEqual(got, []int{2, 3, 4}) {
		t.Errorf("Items() = %v, want [2 3 4]", got)
	}
	r.Add(5)
	if got := r.Items(); !reflect.DeepEqual(got, []int{3, 4, 5}) {
		t.Errorf("Items() = %v, want [3 4 5]", got)
	}
}

func TestRingFillsGradually(t *testing.T) {
	r := NewRing[int](3)
	r.Add(1)
	if got := r.Items(); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Items() = %v, want [1]", got)
	}
	r.Add(2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Items() = %v, want [1 2]", got)
	}
}

func TestRingEmpty(t *testing.T) {
	r := NewRing[int](2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Items() = %v, want []", got)
	}
}
