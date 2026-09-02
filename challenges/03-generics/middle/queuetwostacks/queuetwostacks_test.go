package queuetwostacks

import "testing"

func TestSQueueFIFO(t *testing.T) {
	var q SQueue[int]
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	for _, w := range []int{1, 2, 3} {
		got, ok := q.Dequeue()
		if !ok || got != w {
			t.Fatalf("Dequeue() = %v, %v, want %v, true", got, ok, w)
		}
	}
}

func TestSQueueInterleaved(t *testing.T) {
	var q SQueue[int]
	q.Enqueue(1)
	if v, _ := q.Dequeue(); v != 1 {
		t.Fatalf("Dequeue() = %v, want 1", v)
	}
	q.Enqueue(2)
	q.Enqueue(3)
	if v, _ := q.Dequeue(); v != 2 {
		t.Errorf("Dequeue() = %v, want 2", v)
	}
	if q.Len() != 1 {
		t.Errorf("Len() = %d, want 1", q.Len())
	}
}

func TestSQueueEmpty(t *testing.T) {
	var q SQueue[int]
	if v, ok := q.Dequeue(); v != 0 || ok {
		t.Errorf("Dequeue() on empty = %v, %v, want 0, false", v, ok)
	}
	if q.Len() != 0 {
		t.Errorf("Len() = %d, want 0", q.Len())
	}
}
