package asinjoin

import (
	"errors"
	"fmt"
	"testing"
)

func TestFirstField(t *testing.T) {
	age := &FieldError{Name: "age"}
	email := &FieldError{Name: "email"}

	t.Run("nil", func(t *testing.T) {
		if got, ok := FirstField(nil); ok || got != nil {
			t.Errorf("FirstField(nil) = %v, %v, want nil, false", got, ok)
		}
	})

	t.Run("direct", func(t *testing.T) {
		got, ok := FirstField(age)
		if !ok || got != age {
			t.Errorf("FirstField = %v, %v, want the *FieldError, true", got, ok)
		}
	})

	t.Run("inside_join", func(t *testing.T) {
		got, ok := FirstField(errors.Join(ErrOther, age))
		if !ok || got != age {
			t.Errorf("FirstField = %v, %v, want the *FieldError, true", got, ok)
		}
	})

	t.Run("first_of_two", func(t *testing.T) {
		got, ok := FirstField(errors.Join(email, age))
		if !ok || got != email {
			t.Errorf("FirstField = %v, want the first *FieldError %v", got, email)
		}
	})

	t.Run("wrapped_join", func(t *testing.T) {
		got, ok := FirstField(fmt.Errorf("validate: %w", errors.Join(ErrOther, age)))
		if !ok || got != age {
			t.Errorf("FirstField = %v, %v, want the *FieldError, true", got, ok)
		}
	})

	t.Run("absent", func(t *testing.T) {
		if got, ok := FirstField(errors.Join(ErrOther, errors.New("x"))); ok || got != nil {
			t.Errorf("FirstField = %v, %v, want nil, false", got, ok)
		}
	})
}
