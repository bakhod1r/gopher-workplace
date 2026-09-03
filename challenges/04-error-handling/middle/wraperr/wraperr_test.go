package wraperr

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	t.Run("nil_passthrough", func(t *testing.T) {
		if got := Wrap("read", nil); got != nil {
			t.Errorf("Wrap(\"read\", nil) = %v, want nil", got)
		}
	})

	t.Run("message", func(t *testing.T) {
		got := Wrap("read", ErrDisk)
		if got == nil {
			t.Fatal("Wrap returned nil, want error")
		}
		if got.Error() != "read: disk offline" {
			t.Errorf("message = %q, want %q", got.Error(), "read: disk offline")
		}
	})

	t.Run("preserves_identity", func(t *testing.T) {
		got := Wrap("read", ErrDisk)
		if !errors.Is(got, ErrDisk) {
			t.Errorf("errors.Is(%v, ErrDisk) = false, want true", got)
		}
	})

	t.Run("unwraps_to_original", func(t *testing.T) {
		got := errors.Unwrap(Wrap("read", ErrDisk))
		if got != ErrDisk {
			t.Errorf("Unwrap = %v, want ErrDisk", got)
		}
	})

	t.Run("nested", func(t *testing.T) {
		got := Wrap("handler", Wrap("read", ErrDisk))
		if got.Error() != "handler: read: disk offline" {
			t.Errorf("message = %q, want %q", got.Error(), "handler: read: disk offline")
		}
		if !errors.Is(got, ErrDisk) {
			t.Error("nested wrap lost the sentinel")
		}
	})
}
