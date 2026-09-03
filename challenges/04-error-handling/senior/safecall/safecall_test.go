package safecall

import (
	"errors"
	"testing"
)

func TestSafeCall(t *testing.T) {
	errStep := errors.New("step failed")

	t.Run("nil_func", func(t *testing.T) {
		if err := SafeCall(nil); !errors.Is(err, ErrNilFunc) {
			t.Errorf("err = %v, want ErrNilFunc", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		if err := SafeCall(func() error { return nil }); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("normal_error", func(t *testing.T) {
		err := SafeCall(func() error { return errStep })
		if !errors.Is(err, errStep) {
			t.Errorf("err = %v, want errStep", err)
		}
	})

	t.Run("string_panic", func(t *testing.T) {
		err := SafeCall(func() error { panic("boom") })
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		if err.Error() != "panic: boom" {
			t.Errorf("message = %q, want %q", err.Error(), "panic: boom")
		}
	})

	t.Run("int_panic", func(t *testing.T) {
		err := SafeCall(func() error { panic(7) })
		if err == nil || err.Error() != "panic: 7" {
			t.Errorf("message = %v, want %q", err, "panic: 7")
		}
	})
}
