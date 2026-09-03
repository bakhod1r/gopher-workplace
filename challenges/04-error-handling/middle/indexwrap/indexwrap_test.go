package indexwrap

import (
	"errors"
	"testing"
)

func TestAtIndex(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := AtIndex(0, nil); got != nil {
			t.Errorf("AtIndex(0, nil) = %v, want nil", got)
		}
	})

	t.Run("message", func(t *testing.T) {
		got := AtIndex(3, ErrParse)
		if got == nil {
			t.Fatal("AtIndex returned nil")
		}
		if got.Error() != "record 3: parse failed" {
			t.Errorf("message = %q, want %q", got.Error(), "record 3: parse failed")
		}
	})

	t.Run("zero_index", func(t *testing.T) {
		got := AtIndex(0, ErrParse)
		if got.Error() != "record 0: parse failed" {
			t.Errorf("message = %q, want %q", got.Error(), "record 0: parse failed")
		}
	})

	t.Run("matchable", func(t *testing.T) {
		if !errors.Is(AtIndex(7, ErrParse), ErrParse) {
			t.Error("errors.Is = false, want true")
		}
	})
}
