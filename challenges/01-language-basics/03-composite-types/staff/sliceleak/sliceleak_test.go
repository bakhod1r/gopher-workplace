package sliceleak

import "testing"

func TestHead3(t *testing.T) {
	xs := make([]int, 3, 100000) // huge backing capacity
	xs[0], xs[1], xs[2] = 1, 2, 3
	h := Head3(xs)
	if len(h) != 3 || h[0] != 1 || h[2] != 3 {
		t.Fatalf("contents wrong: %v", h)
	}
	if cap(h) > 3 {
		t.Errorf("retains backing array: cap=%d; want <= 3", cap(h))
	}
}
