package findsentinel

import (
	"fmt"
	"testing"
)

func TestFirstMatch(t *testing.T) {
	wrapped := fmt.Errorf("replica 2: %w", ErrTimeout)

	t.Run("nil_slice", func(t *testing.T) {
		if got := FirstMatch(nil, ErrTimeout); got != nil {
			t.Errorf("FirstMatch = %v, want nil", got)
		}
	})

	t.Run("no_match", func(t *testing.T) {
		if got := FirstMatch([]error{ErrOther, nil}, ErrTimeout); got != nil {
			t.Errorf("FirstMatch = %v, want nil", got)
		}
	})

	t.Run("returns_wrapper", func(t *testing.T) {
		got := FirstMatch([]error{nil, ErrOther, wrapped}, ErrTimeout)
		if got != wrapped {
			t.Errorf("FirstMatch = %v, want the wrapper %v", got, wrapped)
		}
	})

	t.Run("first_of_two", func(t *testing.T) {
		second := fmt.Errorf("replica 3: %w", ErrTimeout)
		got := FirstMatch([]error{wrapped, second}, ErrTimeout)
		if got != wrapped {
			t.Errorf("FirstMatch = %v, want the first match", got)
		}
	})
}
