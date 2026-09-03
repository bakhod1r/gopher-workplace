package causeerr

import (
	"errors"
	"fmt"
	"testing"
)

func TestCodeError(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		err := &CodeError{Code: 7, Cause: ErrDB}
		if err.Error() != "[7] db unavailable" {
			t.Errorf("Error() = %q, want %q", err.Error(), "[7] db unavailable")
		}
	})

	t.Run("unwrap", func(t *testing.T) {
		if got := errors.Unwrap(&CodeError{Code: 7, Cause: ErrDB}); got != ErrDB {
			t.Errorf("Unwrap = %v, want ErrDB", got)
		}
	})

	t.Run("matchable", func(t *testing.T) {
		var err error = &CodeError{Code: 7, Cause: ErrDB}
		if !errors.Is(err, ErrDB) {
			t.Error("errors.Is = false, want true")
		}
	})

	t.Run("matchable_when_wrapped", func(t *testing.T) {
		err := fmt.Errorf("handler: %w", &CodeError{Code: 7, Cause: ErrDB})
		if !errors.Is(err, ErrDB) {
			t.Error("errors.Is through fmt wrapper = false, want true")
		}
	})

	t.Run("as_recovers_code", func(t *testing.T) {
		err := fmt.Errorf("handler: %w", &CodeError{Code: 9, Cause: ErrDB})
		var ce *CodeError
		if !errors.As(err, &ce) {
			t.Fatal("errors.As failed")
		}
		if ce.Code != 9 {
			t.Errorf("Code = %d, want 9", ce.Code)
		}
	})
}
