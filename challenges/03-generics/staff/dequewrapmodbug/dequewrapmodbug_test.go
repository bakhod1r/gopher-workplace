package dequewrapmodbug

import (
	"testing"
)

func TestDequeAsQueue(t *testing.T) {
	var d Deque[int]
	d.PushBack(1)
	d.PushBack(2)
	for _, want := range []int{1, 2} {
		got, ok := d.PopFront()
		if !ok || got != want {
			t.Fatalf("PopFront = %d, %v, want %d, true", got, ok, want)
		}
	}
}

func TestDequePushFrontWraps(t *testing.T) {
	var d Deque[int]
	d.PushBack(2)
	d.PushFront(1)
	if d.Len() != 2 {
		t.Fatalf("Len = %d, want 2", d.Len())
	}
	for i, want := range []int{1, 2} {
		got, ok := d.At(i)
		if !ok || got != want {
			t.Fatalf("At(%d) = %d, %v, want %d, true", i, got, ok, want)
		}
	}
}

func TestDequeAlternating(t *testing.T) {
	var d Deque[int]
	for i := 1; i <= 8; i++ {
		if i%2 == 0 {
			d.PushBack(i)
		} else {
			d.PushFront(i)
		}
	}
	want := []int{7, 5, 3, 1, 2, 4, 6, 8}
	for i, w := range want {
		got, ok := d.At(i)
		if !ok || got != w {
			t.Fatalf("At(%d) = %d, %v, want %d, true", i, got, ok, w)
		}
	}
	for _, w := range want {
		got, ok := d.PopFront()
		if !ok || got != w {
			t.Fatalf("PopFront = %d, %v, want %d, true", got, ok, w)
		}
	}
	if _, ok := d.PopFront(); ok {
		t.Error("PopFront on empty deque = ok, want false")
	}
}
