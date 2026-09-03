package binheapchildbug

import (
	"testing"
	"time"
)

func drain(t *testing.T, h *Heap[int], n int) []int {
	t.Helper()
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		v, ok := h.Pop()
		if !ok {
			t.Fatalf("Pop %d: heap empty early", i)
		}
		out = append(out, v)
	}
	return out
}

func TestHeapSmall(t *testing.T) {
	var h Heap[int]
	for _, v := range []int{5, 3, 8, 1} {
		h.Push(v)
	}
	got := drain(t, &h, 4)
	want := []int{1, 3, 5, 8}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("drain = %v, want %v", got, want)
		}
	}
}

func TestHeapDuplicates(t *testing.T) {
	var h Heap[int]
	for _, v := range []int{2, 2, 1, 2, 1} {
		h.Push(v)
	}
	got := drain(t, &h, 5)
	want := []int{1, 1, 2, 2, 2}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("drain = %v, want %v", got, want)
		}
	}
}

func TestHeapScale(t *testing.T) {
	const n = 10000
	start := time.Now()
	var h Heap[int]
	x := uint32(12345)
	for i := 0; i < n; i++ {
		x = x*1664525 + 1013904223
		h.Push(int(x >> 8))
	}
	if h.Len() != n {
		t.Fatalf("Len = %d, want %d", h.Len(), n)
	}
	prev := -1
	for i := 0; i < n; i++ {
		v, ok := h.Pop()
		if !ok {
			t.Fatalf("Pop %d: heap empty early", i)
		}
		if v < prev {
			t.Fatalf("Pop %d = %d, which is less than the previous %d", i, v, prev)
		}
		prev = v
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("10k push/pop took %v, want under 2s", d)
	}
}
