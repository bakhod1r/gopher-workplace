package ringbufwriter

import (
	"reflect"
	"testing"
)

func TestRingKeepsTheMostRecent(t *testing.T) {
	r := New(3)
	for _, v := range []int{1, 2, 3, 4} {
		r.Add(v)
	}
	if got := r.Snapshot(); !reflect.DeepEqual(got, []int{2, 3, 4}) {
		t.Errorf("Snapshot = %v, want [2 3 4]", got)
	}
	if r.Len() != 3 {
		t.Errorf("Len = %d, want 3", r.Len())
	}
	if r.Total() != 4 {
		t.Errorf("Total = %d, want 4", r.Total())
	}
}

func TestRingPartiallyFilled(t *testing.T) {
	r := New(4)
	r.Add(1)
	r.Add(2)
	if got := r.Snapshot(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Snapshot = %v, want [1 2]", got)
	}
	if r.Len() != 2 {
		t.Errorf("Len = %d, want 2", r.Len())
	}
}

func TestRingEmpty(t *testing.T) {
	r := New(3)
	got := r.Snapshot()
	if got == nil || len(got) != 0 {
		t.Errorf("Snapshot = %v, want empty non-nil slice", got)
	}
	if r.Len() != 0 || r.Total() != 0 {
		t.Errorf("Len, Total = %d, %d, want 0, 0", r.Len(), r.Total())
	}
}

func TestRingZeroCapacity(t *testing.T) {
	r := New(0)
	r.Add(1)
	r.Add(2)
	if got := r.Snapshot(); len(got) != 0 {
		t.Errorf("Snapshot = %v, want empty", got)
	}
	if r.Total() != 2 {
		t.Errorf("Total = %d, want 2 — adds still count", r.Total())
	}
}

func TestRingWrapsRepeatedly(t *testing.T) {
	r := New(3)
	for i := 1; i <= 100; i++ {
		r.Add(i)
	}
	if got := r.Snapshot(); !reflect.DeepEqual(got, []int{98, 99, 100}) {
		t.Errorf("Snapshot = %v, want [98 99 100]", got)
	}
}

func TestAddDoesNotAllocate(t *testing.T) {
	r := New(64)
	for i := 0; i < 64; i++ {
		r.Add(i)
	}
	if allocs := testing.AllocsPerRun(100, func() { r.Add(7) }); allocs != 0 {
		t.Errorf("Add made %v allocations, want 0", allocs)
	}
}

func TestSnapshotIsACopy(t *testing.T) {
	r := New(3)
	r.Add(1)
	r.Add(2)
	s := r.Snapshot()
	s[0] = 99
	if got := r.Snapshot(); got[0] != 1 {
		t.Errorf("Snapshot aliased the ring: %v", got)
	}
}
