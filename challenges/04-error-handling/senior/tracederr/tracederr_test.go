package tracederr

import (
	"errors"
	"fmt"
	"testing"
)

func TestTracedError(t *testing.T) {
	te := &TracedError{Op: "load", Cause: ErrDisk}

	t.Run("short_message", func(t *testing.T) {
		if te.Error() != "disk offline" {
			t.Errorf("Error() = %q, want %q", te.Error(), "disk offline")
		}
	})

	t.Run("trace", func(t *testing.T) {
		if te.Trace() != "load -> disk offline" {
			t.Errorf("Trace() = %q, want %q", te.Trace(), "load -> disk offline")
		}
	})

	t.Run("unwrap", func(t *testing.T) {
		if errors.Unwrap(te) != ErrDisk {
			t.Errorf("Unwrap = %v, want ErrDisk", errors.Unwrap(te))
		}
	})

	t.Run("matchable_when_wrapped", func(t *testing.T) {
		err := fmt.Errorf("handler: %w", te)
		if !errors.Is(err, ErrDisk) {
			t.Error("errors.Is = false, want true")
		}
	})
}
