package ringbuf

import (
	"reflect"
	"testing"
)

func TestRingFillsThenOverwrites(t *testing.T) {
	r := NewRing(3)
	for _, v := range []int{1, 2, 3} {
		r.Push(v)
	}
	if got := r.Items(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("Items = %v, want [1 2 3]", got)
	}
	r.Push(4)
	if got := r.Items(); !reflect.DeepEqual(got, []int{2, 3, 4}) {
		t.Errorf("Items = %v, want [2 3 4]", got)
	}
	if r.Len() != 3 {
		t.Errorf("Len = %d, want 3", r.Len())
	}
}

func TestRingCapacityOne(t *testing.T) {
	r := NewRing(1)
	r.Push(1)
	r.Push(2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Items = %v, want [2]", got)
	}
}

func TestPushNeverAllocates(t *testing.T) {
	r := NewRing(8)
	if n := testing.AllocsPerRun(1000, func() { r.Push(1) }); n != 0 {
		t.Errorf("Push made %v allocations, want 0: the ring must not grow", n)
	}
}
