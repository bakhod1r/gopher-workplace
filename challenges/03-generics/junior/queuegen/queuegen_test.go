package queuegen

import "testing"

func TestQueue(t *testing.T) {
	var q Queue[int]
	if _, ok := q.Dequeue(); ok {
		t.Error("Dequeue() on empty queue reported ok, want false")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	if q.Len() != 2 {
		t.Errorf("Len() = %d, want 2", q.Len())
	}
	if v, ok := q.Dequeue(); v != 1 || !ok {
		t.Errorf("Dequeue() = %v, %v, want 1, true (FIFO)", v, ok)
	}
	if v, ok := q.Dequeue(); v != 2 || !ok {
		t.Errorf("Dequeue() = %v, %v, want 2, true", v, ok)
	}
	if q.Len() != 0 {
		t.Errorf("Len() = %d, want 0", q.Len())
	}
}
