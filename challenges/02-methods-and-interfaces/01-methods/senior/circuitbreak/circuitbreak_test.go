package circuitbreak

import (
	"errors"
	"testing"
)

func TestBreaker(t *testing.T) {
	b := &Breaker{Threshold: 2}
	failFn := func() error { return errors.New("fail") }
	okFn := func() error { return nil }

	if err := b.Call(failFn); err == nil {
		t.Error("expected error")
	}
	if b.IsOpen {
		t.Error("should not be open yet")
	}

	if err := b.Call(failFn); err == nil {
		t.Error("expected error")
	}
	if !b.IsOpen {
		t.Error("should be open now")
	}

	err := b.Call(okFn)
	if err == nil || err.Error() != "circuit open" {
		t.Errorf("expected circuit open error, got %v", err)
	}
}
