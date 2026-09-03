package validationerr

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidationError(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		err := NewValidation("email", "is required")
		if err.Error() != "email: is required" {
			t.Errorf("Error() = %q, want %q", err.Error(), "email: is required")
		}
	})

	t.Run("fields_readable", func(t *testing.T) {
		var ve *ValidationError
		if !errors.As(NewValidation("age", "must be positive"), &ve) {
			t.Fatal("errors.As failed, want *ValidationError")
		}
		if ve.Field != "age" || ve.Reason != "must be positive" {
			t.Errorf("fields = %q, %q, want %q, %q", ve.Field, ve.Reason, "age", "must be positive")
		}
	})

	t.Run("survives_wrapping", func(t *testing.T) {
		wrapped := fmt.Errorf("save: %w", NewValidation("email", "is required"))
		var ve *ValidationError
		if !errors.As(wrapped, &ve) {
			t.Fatal("errors.As through wrapping failed")
		}
		if ve.Field != "email" {
			t.Errorf("Field = %q, want %q", ve.Field, "email")
		}
	})
}
