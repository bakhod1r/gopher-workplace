package safeis

import (
	"fmt"
	"testing"
)

// uncomparable has a slice field, so == on it panics.
type uncomparable struct {
	parts []string
}

func (e uncomparable) Error() string { return "uncomparable" }

func TestSafeIs(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		if !SafeIs(ErrA, ErrA) {
			t.Error("SafeIs = false, want true")
		}
	})

	t.Run("no_match", func(t *testing.T) {
		if SafeIs(ErrA, fmt.Errorf("other")) {
			t.Error("SafeIs = true, want false")
		}
	})

	t.Run("wrapped_match", func(t *testing.T) {
		if !SafeIs(fmt.Errorf("x: %w", ErrA), ErrA) {
			t.Error("SafeIs = false, want true")
		}
	})

	t.Run("nil_arguments", func(t *testing.T) {
		if SafeIs(nil, ErrA) || SafeIs(ErrA, nil) {
			t.Error("SafeIs with a nil argument = true, want false")
		}
	})

	t.Run("uncomparable_does_not_panic", func(t *testing.T) {
		u := uncomparable{parts: []string{"x"}}
		if got := SafeIs(u, u); got {
			t.Errorf("SafeIs = %v, want false", got)
		}
	})
}
