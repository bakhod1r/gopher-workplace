package heapgen

import "testing"

func TestHeapOrder(t *testing.T) {
	var h Heap[int]
	for _, v := range []int{5, 3, 8, 1, 9, 2} {
		h.Push(v)
	}
	if h.Len() != 6 {
		t.Fatalf("Len() = %d, want 6", h.Len())
	}
	want := []int{1, 2, 3, 5, 8, 9}
	for _, w := range want {
		got, ok := h.Pop()
		if !ok || got != w {
			t.Fatalf("Pop() = %v, %v, want %v, true", got, ok, w)
		}
	}
	if h.Len() != 0 {
		t.Errorf("Len() = %d, want 0", h.Len())
	}
}

func TestHeapEmpty(t *testing.T) {
	var h Heap[int]
	if v, ok := h.Pop(); v != 0 || ok {
		t.Errorf("Pop() on empty = %v, %v, want 0, false", v, ok)
	}
}

func TestHeapStrings(t *testing.T) {
	var h Heap[string]
	h.Push("b")
	h.Push("a")
	if v, _ := h.Pop(); v != "a" {
		t.Errorf("Pop() = %q, want a", v)
	}
}
