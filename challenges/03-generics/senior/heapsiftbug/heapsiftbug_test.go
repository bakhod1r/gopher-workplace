package heapsiftbug

import "testing"

func TestHeapDrainsInOrder(t *testing.T) {
	var h Heap[int]
	for _, v := range []int{5, 3, 8, 1, 9, 2, 7, 4, 6} {
		h.Push(v)
	}
	prev := -1
	for h.Len() > 0 {
		got, ok := h.Pop()
		if !ok {
			t.Fatal("Pop reported empty too early")
		}
		if got < prev {
			t.Fatalf("Pop returned %d after %d: the heap order is broken", got, prev)
		}
		prev = got
	}
}

func TestHeapRightChildSmaller(t *testing.T) {
	var h Heap[int]
	for _, v := range []int{10, 8, 3, 9, 7} {
		h.Push(v)
	}
	want := []int{3, 7, 8, 9, 10}
	for _, w := range want {
		got, _ := h.Pop()
		if got != w {
			t.Fatalf("Pop = %d, want %d", got, w)
		}
	}
}

func TestHeapEmpty(t *testing.T) {
	var h Heap[int]
	if v, ok := h.Pop(); v != 0 || ok {
		t.Errorf("Pop on empty = %v, %v, want 0, false", v, ok)
	}
}
