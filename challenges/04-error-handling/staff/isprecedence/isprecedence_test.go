package isprecedence

import (
	"errors"
	"fmt"
	"testing"
)

func TestCodedError(t *testing.T) {
	e := &CodedError{Code: 7, Cause: ErrBase}

	t.Run("message", func(t *testing.T) {
		if e.Error() != "code 7: base" {
			t.Errorf("Error() = %q, want %q", e.Error(), "code 7: base")
		}
	})

	t.Run("matches_by_code", func(t *testing.T) {
		if !errors.Is(e, &CodedError{Code: 7}) {
			t.Error("errors.Is by code = false, want true")
		}
	})

	t.Run("rejects_other_code", func(t *testing.T) {
		if errors.Is(e, &CodedError{Code: 9}) {
			t.Error("errors.Is matched a different code")
		}
	})

	t.Run("cause_still_matches", func(t *testing.T) {
		if !errors.Is(e, ErrBase) {
			t.Error("errors.Is for the cause = false, want true")
		}
	})

	t.Run("through_wrapping", func(t *testing.T) {
		wrapped := fmt.Errorf("layer: %w", e)
		if !errors.Is(wrapped, ErrBase) || !errors.Is(wrapped, &CodedError{Code: 7}) {
			t.Error("wrapped error lost a match")
		}
	})
}
