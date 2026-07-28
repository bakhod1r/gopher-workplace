package minseed

import "testing"

func TestMinPtr(t *testing.T) {
	xs := []int{1, 2, 1}
	p := MinPtr(xs)
	if *p != 1 {
		t.Fatalf("*p=%d want 1", *p)
	}
	if p != &xs[0] {
		t.Fatalf("should point to the FIRST minimum (index 0), seeding from the end returns the last")
	}
}
