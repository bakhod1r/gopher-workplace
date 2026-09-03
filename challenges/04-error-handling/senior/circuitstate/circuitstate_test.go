package circuitstate

import (
	"errors"
	"testing"
)

func TestBreaker(t *testing.T) {
	errBoom := errors.New("boom")

	t.Run("success", func(t *testing.T) {
		b := &Breaker{Threshold: 2}
		if err := b.Call(func() error { return nil }); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("opens_after_threshold", func(t *testing.T) {
		b := &Breaker{Threshold: 2}
		calls := 0
		fail := func() error { calls++; return errBoom }

		if err := b.Call(fail); !errors.Is(err, errBoom) {
			t.Fatalf("first call err = %v, want errBoom", err)
		}
		if err := b.Call(fail); !errors.Is(err, errBoom) {
			t.Fatalf("second call err = %v, want errBoom", err)
		}
		if err := b.Call(fail); !errors.Is(err, ErrOpen) {
			t.Fatalf("third call err = %v, want ErrOpen", err)
		}
		if calls != 2 {
			t.Errorf("calls = %d, want 2 (the third must not run)", calls)
		}
	})

	t.Run("success_resets", func(t *testing.T) {
		b := &Breaker{Threshold: 2}
		b.Call(func() error { return errBoom })
		b.Call(func() error { return nil })
		if err := b.Call(func() error { return errBoom }); !errors.Is(err, errBoom) {
			t.Errorf("err = %v, want errBoom — the success should have reset the count", err)
		}
	})

	t.Run("zero_threshold", func(t *testing.T) {
		b := &Breaker{Threshold: 0}
		calls := 0
		err := b.Call(func() error { calls++; return nil })
		if !errors.Is(err, ErrOpen) || calls != 0 {
			t.Errorf("err = %v, calls = %d, want ErrOpen and 0 calls", err, calls)
		}
	})
}
