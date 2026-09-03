package asall

import (
	"errors"
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	a := &FieldError{Name: "a"}
	b := &FieldError{Name: "b"}

	t.Run("nil", func(t *testing.T) {
		if got := All(nil); got != nil {
			t.Errorf("All(nil) = %v, want nil", got)
		}
	})

	t.Run("none", func(t *testing.T) {
		if got := All(ErrOther); got != nil {
			t.Errorf("All = %v, want nil", got)
		}
	})

	t.Run("direct", func(t *testing.T) {
		got := All(a)
		if len(got) != 1 || got[0] != a {
			t.Errorf("All = %v, want [a]", got)
		}
	})

	t.Run("joined", func(t *testing.T) {
		got := All(errors.Join(a, ErrOther, b))
		if len(got) != 2 || got[0] != a || got[1] != b {
			t.Errorf("All = %v, want [a b]", got)
		}
	})

	t.Run("nested", func(t *testing.T) {
		got := All(errors.Join(fmt.Errorf("x: %w", a), errors.Join(b)))
		if len(got) != 2 || got[0] != a || got[1] != b {
			t.Errorf("All = %v, want [a b]", got)
		}
	})
}
