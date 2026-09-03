package panicerror

import (
	"errors"
	"testing"
)

func TestCapture(t *testing.T) {
	t.Run("no_panic", func(t *testing.T) {
		if err := Capture(func() {}); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("error_payload_identity", func(t *testing.T) {
		err := Capture(func() { panic(ErrStop) })
		if err != error(ErrStop) {
			t.Errorf("err = %v, want the ErrStop value itself", err)
		}
	})

	t.Run("error_payload_matchable", func(t *testing.T) {
		err := Capture(func() { panic(ErrStop) })
		if !errors.Is(err, ErrStop) {
			t.Error("errors.Is = false, want true")
		}
	})

	t.Run("string_payload", func(t *testing.T) {
		err := Capture(func() { panic("bug") })
		if err == nil || err.Error() != "panic: bug" {
			t.Errorf("err = %v, want %q", err, "panic: bug")
		}
	})

	t.Run("struct_payload", func(t *testing.T) {
		err := Capture(func() { panic(42) })
		if err == nil || err.Error() != "panic: 42" {
			t.Errorf("err = %v, want %q", err, "panic: 42")
		}
	})
}
