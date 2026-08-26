package ratelimit

import "testing"

func TestLimiter(t *testing.T) {
	l := NewLimiter(2)

	if !l.Allow() {
		t.Error("expected true")
	}
	if !l.Allow() {
		t.Error("expected true")
	}
	if l.Allow() {
		t.Error("expected false (empty)")
	}

	l.Refill(1)
	if !l.Allow() {
		t.Error("expected true (refilled)")
	}
	if l.Allow() {
		t.Error("expected false (empty again)")
	}
}
