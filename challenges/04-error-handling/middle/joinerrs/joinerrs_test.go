package joinerrs

import (
	"errors"
	"testing"
)

func TestJoinAll(t *testing.T) {
	t.Run("nil_slice", func(t *testing.T) {
		if got := JoinAll(nil); got != nil {
			t.Errorf("JoinAll(nil) = %v, want nil", got)
		}
	})

	t.Run("all_nil", func(t *testing.T) {
		if got := JoinAll([]error{nil, nil}); got != nil {
			t.Errorf("JoinAll = %v, want nil", got)
		}
	})

	t.Run("both", func(t *testing.T) {
		got := JoinAll([]error{ErrA, ErrB})
		if !errors.Is(got, ErrA) || !errors.Is(got, ErrB) {
			t.Errorf("JoinAll = %v, want it to match ErrA and ErrB", got)
		}
	})

	t.Run("skips_nil", func(t *testing.T) {
		got := JoinAll([]error{nil, ErrA, nil})
		if !errors.Is(got, ErrA) {
			t.Fatalf("JoinAll = %v, want it to match ErrA", got)
		}
		if errors.Is(got, ErrB) {
			t.Errorf("JoinAll = %v, want it not to match ErrB", got)
		}
	})
}
