package consterr

import (
	"errors"
	"fmt"
	"testing"
)

func TestConstErrors(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		if ErrClosed.Error() != "closed" {
			t.Errorf("Error() = %q, want %q", ErrClosed.Error(), "closed")
		}
	})

	t.Run("usable_as_error", func(t *testing.T) {
		var err error = ErrBusy
		if err.Error() != "busy" {
			t.Errorf("Error() = %q, want %q", err.Error(), "busy")
		}
	})

	t.Run("distinct", func(t *testing.T) {
		if ErrClosed == ErrBusy {
			t.Error("sentinels compare equal, want distinct")
		}
	})

	t.Run("matchable_when_wrapped", func(t *testing.T) {
		err := fmt.Errorf("write: %w", ErrClosed)
		if !errors.Is(err, ErrClosed) {
			t.Error("errors.Is = false, want true")
		}
		if errors.Is(err, ErrBusy) {
			t.Error("errors.Is matched the wrong sentinel")
		}
	})
}
