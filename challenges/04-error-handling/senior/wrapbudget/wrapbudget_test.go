package wrapbudget

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := Wrap(nil, "a", 2); got != nil {
			t.Errorf("Wrap(nil, ...) = %v, want nil", got)
		}
	})

	t.Run("wraps_below_limit", func(t *testing.T) {
		got := Wrap(ErrBase, "a", 2)
		if got.Error() != "a: base" {
			t.Errorf("message = %q, want %q", got.Error(), "a: base")
		}
	})

	t.Run("refuses_at_limit", func(t *testing.T) {
		two := Wrap(ErrBase, "a", 5)
		if got := Wrap(two, "b", 2); got != two {
			t.Errorf("Wrap = %v, want the input unchanged", got)
		}
	})

	t.Run("still_matchable", func(t *testing.T) {
		got := Wrap(Wrap(ErrBase, "a", 5), "b", 2)
		if !errors.Is(got, ErrBase) {
			t.Error("errors.Is = false, want true")
		}
	})

	t.Run("zero_max", func(t *testing.T) {
		if got := Wrap(ErrBase, "a", 0); got != error(ErrBase) {
			t.Errorf("Wrap = %v, want ErrBase unchanged", got)
		}
	})
}
