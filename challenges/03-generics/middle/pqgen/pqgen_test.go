package pqgen

import "testing"

func TestPQPriority(t *testing.T) {
	var q PQ[string]
	q.Push("low", 5)
	q.Push("high", 1)
	q.Push("mid", 3)
	want := []string{"high", "mid", "low"}
	for _, w := range want {
		got, ok := q.Pop()
		if !ok || got != w {
			t.Fatalf("Pop() = %q, %v, want %q", got, ok, w)
		}
	}
}

func TestPQTiesAreFIFO(t *testing.T) {
	var q PQ[string]
	q.Push("a", 1)
	q.Push("b", 1)
	q.Push("c", 1)
	for _, w := range []string{"a", "b", "c"} {
		got, _ := q.Pop()
		if got != w {
			t.Fatalf("Pop() = %q, want %q (equal priorities are FIFO)", got, w)
		}
	}
}

func TestPQEmpty(t *testing.T) {
	var q PQ[int]
	if v, ok := q.Pop(); v != 0 || ok {
		t.Errorf("Pop() on empty = %v, %v, want 0, false", v, ok)
	}
	q.Push(1, 1)
	if q.Len() != 1 {
		t.Errorf("Len() = %d, want 1", q.Len())
	}
}
