package circuitifc

import (
	"errors"
	"testing"
	"time"
)

var errBoom = errors.New("boom")

func breaker(threshold int) (*Breaker, *FakeClock) {
	c := &FakeClock{T: time.Unix(0, 0)}
	return NewBreaker(threshold, time.Minute, c), c
}

func failing(calls *int) Op {
	return OpFunc(func() error {
		*calls++
		return errBoom
	})
}

func succeeding(calls *int) Op {
	return OpFunc(func() error {
		*calls++
		return nil
	})
}

func TestClosedPassesThrough(t *testing.T) {
	b, _ := breaker(2)
	calls := 0
	if err := b.Call(succeeding(&calls)); err != nil {
		t.Errorf("Call = %v, want nil", err)
	}
	if calls != 1 {
		t.Errorf("calls = %d, want 1", calls)
	}
	if b.IsOpen() {
		t.Error("breaker should stay closed")
	}
}

func TestOpensAfterThreshold(t *testing.T) {
	b, _ := breaker(2)
	calls := 0
	b.Call(failing(&calls))
	if b.IsOpen() {
		t.Error("one failure should not open the breaker")
	}
	b.Call(failing(&calls))
	if !b.IsOpen() {
		t.Error("threshold failures should open the breaker")
	}
}

func TestOpenFailsFastWithoutCalling(t *testing.T) {
	b, _ := breaker(1)
	calls := 0
	b.Call(failing(&calls))
	before := calls

	if err := b.Call(failing(&calls)); !errors.Is(err, ErrOpen) {
		t.Errorf("Call = %v, want ErrOpen", err)
	}
	if calls != before {
		t.Errorf("the operation ran while the breaker was open (%d calls)", calls-before)
	}
}

func TestProbeAfterCooldownCloses(t *testing.T) {
	b, c := breaker(1)
	calls := 0
	b.Call(failing(&calls))

	c.Advance(time.Minute)
	if err := b.Call(succeeding(&calls)); err != nil {
		t.Errorf("probe = %v, want nil", err)
	}
	if b.IsOpen() {
		t.Error("a successful probe should close the breaker")
	}
}

func TestFailedProbeReopens(t *testing.T) {
	b, c := breaker(1)
	calls := 0
	b.Call(failing(&calls))

	c.Advance(time.Minute)
	if err := b.Call(failing(&calls)); !errors.Is(err, errBoom) {
		t.Errorf("probe = %v, want errBoom", err)
	}
	if !b.IsOpen() {
		t.Error("a failed probe should leave the breaker open")
	}

	before := calls
	if err := b.Call(failing(&calls)); !errors.Is(err, ErrOpen) {
		t.Errorf("Call = %v, want ErrOpen", err)
	}
	if calls != before {
		t.Error("the operation ran while open")
	}
}

func TestSuccessResetsFailureCount(t *testing.T) {
	b, _ := breaker(2)
	calls := 0
	b.Call(failing(&calls))
	b.Call(succeeding(&calls))
	b.Call(failing(&calls))
	if b.IsOpen() {
		t.Error("a success should have reset the consecutive failure count")
	}
}
