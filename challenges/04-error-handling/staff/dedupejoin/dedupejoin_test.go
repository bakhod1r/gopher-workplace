package dedupejoin

import (
	"errors"
	"strings"
	"testing"
)

func TestDistinct(t *testing.T) {
	sameAsA := errors.New("a")

	t.Run("nothing", func(t *testing.T) {
		if got := Distinct(); got != nil {
			t.Errorf("Distinct() = %v, want nil", got)
		}
	})

	t.Run("all_nil", func(t *testing.T) {
		if got := Distinct(nil, nil); got != nil {
			t.Errorf("Distinct = %v, want nil", got)
		}
	})

	t.Run("dedupes", func(t *testing.T) {
		got := Distinct(ErrA, sameAsA, ErrB)
		if got == nil {
			t.Fatal("got nil, want an error")
		}
		if lines := strings.Count(got.Error(), "\n") + 1; lines != 2 {
			t.Errorf("message %q has %d lines, want 2", got.Error(), lines)
		}
		if !errors.Is(got, ErrA) || !errors.Is(got, ErrB) {
			t.Error("errors.Is failed for a kept error")
		}
	})

	t.Run("keeps_first_value", func(t *testing.T) {
		got := Distinct(ErrA, sameAsA)
		if !errors.Is(got, ErrA) {
			t.Error("errors.Is(ErrA) = false, want the first occurrence kept")
		}
		if errors.Is(got, sameAsA) {
			t.Error("the later duplicate was kept as well")
		}
	})
}
