package ringbuffer

import "testing"

func TestRingBuffer(t *testing.T) {
	r := New(2)
	if err := r.Push(1); err != nil {
		t.Fatal(err)
	}
	if err := r.Push(2); err != nil {
		t.Fatal(err)
	}
	if err := r.Push(3); err == nil {
		t.Error("expected full error")
	}
}
