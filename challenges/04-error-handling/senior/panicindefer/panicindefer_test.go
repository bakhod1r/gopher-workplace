package panicindefer

import (
	"errors"
	"testing"
)

func TestGuard(t *testing.T) {
	errWork := errors.New("work failed")

	t.Run("both_ok", func(t *testing.T) {
		cleaned := false
		err := Guard(func() error { return nil }, func() { cleaned = true })
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if !cleaned {
			t.Error("cleanup did not run")
		}
	})

	t.Run("work_error", func(t *testing.T) {
		cleaned := false
		err := Guard(func() error { return errWork }, func() { cleaned = true })
		if !errors.Is(err, errWork) {
			t.Fatalf("err = %v, want errWork", err)
		}
		if !cleaned {
			t.Error("cleanup did not run after a work failure")
		}
	})

	t.Run("work_panic", func(t *testing.T) {
		cleaned := false
		err := Guard(func() error { panic("boom") }, func() { cleaned = true })
		if !errors.Is(err, ErrPanic) {
			t.Fatalf("err = %v, want it to match ErrPanic", err)
		}
		if !cleaned {
			t.Error("cleanup did not run after a work panic")
		}
	})

	t.Run("cleanup_panic", func(t *testing.T) {
		err := Guard(func() error { return nil }, func() { panic("cleanup boom") })
		if !errors.Is(err, ErrPanic) {
			t.Errorf("err = %v, want it to match ErrPanic", err)
		}
	})
}
