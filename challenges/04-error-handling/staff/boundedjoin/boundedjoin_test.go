package boundedjoin

import (
	"errors"
	"strings"
	"testing"
)

func TestCap(t *testing.T) {
	t.Run("nothing", func(t *testing.T) {
		if got := Cap(2); got != nil {
			t.Errorf("Cap(2) = %v, want nil", got)
		}
	})

	t.Run("all_nil", func(t *testing.T) {
		if got := Cap(2, nil, nil); got != nil {
			t.Errorf("Cap = %v, want nil", got)
		}
	})

	t.Run("under_cap", func(t *testing.T) {
		got := Cap(5, ErrA)
		if !errors.Is(got, ErrA) {
			t.Fatalf("got = %v, want it to match ErrA", got)
		}
		if strings.Contains(got.Error(), "more") {
			t.Errorf("message = %q, want no summary line", got.Error())
		}
	})

	t.Run("over_cap", func(t *testing.T) {
		got := Cap(2, ErrA, ErrB, ErrC)
		if !errors.Is(got, ErrA) || !errors.Is(got, ErrB) {
			t.Errorf("got = %v, want the first two kept", got)
		}
		if errors.Is(got, ErrC) {
			t.Error("an error past the cap was kept")
		}
		if !strings.Contains(got.Error(), "and 1 more") {
			t.Errorf("message = %q, want it to contain %q", got.Error(), "and 1 more")
		}
	})

	t.Run("skips_nil_when_counting", func(t *testing.T) {
		got := Cap(1, nil, ErrA, nil, ErrB)
		if !strings.Contains(got.Error(), "and 1 more") {
			t.Errorf("message = %q, want it to report exactly one dropped error", got.Error())
		}
	})
}
