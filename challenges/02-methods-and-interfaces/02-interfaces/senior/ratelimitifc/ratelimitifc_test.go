package ratelimitifc

import (
	"testing"
	"time"
)

func bucket(burst int) (*TokenBucket, *FakeClock) {
	c := &FakeClock{T: time.Unix(0, 0)}
	return NewTokenBucket(burst, time.Second, c), c
}

func TestBurstThenDeny(t *testing.T) {
	b, _ := bucket(2)
	if !b.Allow() || !b.Allow() {
		t.Fatal("the first two calls should be allowed")
	}
	if b.Allow() {
		t.Error("the third call should be denied")
	}
}

func TestRefill(t *testing.T) {
	b, c := bucket(2)
	b.AllowN(2)

	c.Advance(time.Second)
	if !b.Allow() {
		t.Error("one token should have refilled")
	}
	if b.Allow() {
		t.Error("only one token should have refilled")
	}
}

func TestRefillCappedAtBurst(t *testing.T) {
	b, c := bucket(2)
	b.AllowN(2)

	c.Advance(time.Hour)
	if got := b.AllowN(5); got != 2 {
		t.Errorf("AllowN = %d, want 2 (refill must cap at Burst)", got)
	}
}

func TestAllowN(t *testing.T) {
	b, _ := bucket(2)
	if got := b.AllowN(5); got != 2 {
		t.Errorf("AllowN = %d, want 2", got)
	}
	if got := b.AllowN(0); got != 0 {
		t.Errorf("AllowN(0) = %d, want 0", got)
	}
}

func TestPartialIntervalGivesNothing(t *testing.T) {
	b, c := bucket(1)
	b.Allow()
	c.Advance(999 * time.Millisecond)
	if b.Allow() {
		t.Error("a partial interval must not grant a token")
	}
	c.Advance(time.Millisecond)
	if !b.Allow() {
		t.Error("the full interval should grant a token")
	}
}

func TestIsLimiter(t *testing.T) {
	b, _ := bucket(1)
	var l Limiter = b
	if !l.Allow() {
		t.Error("Allow = false")
	}
}
