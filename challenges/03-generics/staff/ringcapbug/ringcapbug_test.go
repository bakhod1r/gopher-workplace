package ringcapbug

import (
	"reflect"
	"testing"
)

func TestRingSliceOrder(t *testing.T) {
	r := NewRing[int](4)
	for _, v := range []int{1, 2, 3, 4, 5} {
		r.Push(v)
	}
	if got := r.Slice(); !reflect.DeepEqual(got, []int{2, 3, 4, 5}) {
		t.Errorf("Slice = %v, want [2 3 4 5]", got)
	}
	if r.Len() != 4 {
		t.Errorf("Len = %d, want 4", r.Len())
	}
}

func TestRingSliceEmpty(t *testing.T) {
	r := NewRing[string](3)
	if got := r.Slice(); len(got) != 0 {
		t.Errorf("Slice = %v, want []", got)
	}
}

func TestRingSliceCapacityIsClipped(t *testing.T) {
	r := NewRing[int](8)
	r.Push(1)
	r.Push(2)
	r.Push(3)
	s := r.Slice()
	if !reflect.DeepEqual(s, []int{1, 2, 3}) {
		t.Fatalf("Slice = %v, want [1 2 3]", s)
	}
	if cap(s) != len(s) {
		t.Errorf("cap(Slice()) = %d, want %d (len)", cap(s), len(s))
	}
}

func TestRingSliceDoesNotShareSpareCapacity(t *testing.T) {
	r := NewRing[int](8)
	r.Push(1)
	r.Push(2)
	r.Push(3)
	snap := append(r.Slice(), 99)
	r.Push(4)
	if snap[3] != 99 {
		t.Errorf("snapshot[3] = %d, want 99: the ring wrote through the snapshot", snap[3])
	}
	if got := r.Slice(); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("Slice = %v, want [1 2 3 4]", got)
	}
}
