package queuetwostacksbug

import "testing"

func TestQueueInterleaved(t *testing.T) {
	var q Queue[int]
	q.Enqueue(1)
	q.Enqueue(2)
	if got, _ := q.Dequeue(); got != 1 {
		t.Fatalf("Dequeue = %d, want 1", got)
	}
	q.Enqueue(3)
	if got, _ := q.Dequeue(); got != 2 {
		t.Errorf("Dequeue = %d, want 2", got)
	}
	if got, _ := q.Dequeue(); got != 3 {
		t.Errorf("Dequeue = %d, want 3", got)
	}
}

func TestQueueSingle(t *testing.T) {
	var q Queue[int]
	q.Enqueue(1)
	if got, ok := q.Dequeue(); !ok || got != 1 {
		t.Errorf("Dequeue = %d, %v, want 1, true", got, ok)
	}
}

func TestQueueEmpty(t *testing.T) {
	var q Queue[int]
	if got, ok := q.Dequeue(); ok || got != 0 {
		t.Errorf("Dequeue = %d, %v, want 0, false", got, ok)
	}
}
