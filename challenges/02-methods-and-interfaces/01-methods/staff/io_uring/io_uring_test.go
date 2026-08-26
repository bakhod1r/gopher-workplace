package io_uring

import "testing"

func TestRing(t *testing.T) {
	r := &Ring{}
	r.Submit()
	if r.Count != 1 {
		t.Error("expected 1")
	}
}
