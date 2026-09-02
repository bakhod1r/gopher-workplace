package dequegen

import "testing"

func TestDeque(t *testing.T) {
	var d Deque[int]
	if _, ok := d.PopFront(); ok {
		t.Error("PopFront() on empty deque reported ok, want false")
	}
	d.PushBack(1)
	d.PushFront(0)
	d.PushBack(2)
	if v, ok := d.PopFront(); v != 0 || !ok {
		t.Errorf("PopFront() = %v, %v, want 0, true", v, ok)
	}
	if v, ok := d.PopFront(); v != 1 || !ok {
		t.Errorf("PopFront() = %v, %v, want 1, true", v, ok)
	}
	if v, ok := d.PopFront(); v != 2 || !ok {
		t.Errorf("PopFront() = %v, %v, want 2, true", v, ok)
	}
	if _, ok := d.PopFront(); ok {
		t.Error("PopFront() on drained deque reported ok, want false")
	}
}
