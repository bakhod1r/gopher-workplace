package multiwrap

import (
	"errors"
	"testing"
)

func TestBoth(t *testing.T) {
	t.Run("both_nil", func(t *testing.T) {
		if got := Both(nil, nil); got != nil {
			t.Errorf("Both(nil, nil) = %v, want nil", got)
		}
	})

	t.Run("message", func(t *testing.T) {
		got := Both(ErrA, ErrB)
		if got == nil {
			t.Fatal("got nil, want an error")
		}
		if got.Error() != "a; b" {
			t.Errorf("message = %q, want %q", got.Error(), "a; b")
		}
	})

	t.Run("matches_both", func(t *testing.T) {
		got := Both(ErrA, ErrB)
		if !errors.Is(got, ErrA) || !errors.Is(got, ErrB) {
			t.Errorf("errors.Is failed for %v", got)
		}
	})

	t.Run("multi_unwrap_shape", func(t *testing.T) {
		got := Both(ErrA, ErrB)
		joined, ok := got.(interface{ Unwrap() []error })
		if !ok {
			t.Fatal("result does not implement Unwrap() []error")
		}
		if len(joined.Unwrap()) != 2 {
			t.Errorf("unwrapped %d errors, want 2", len(joined.Unwrap()))
		}
	})
}
