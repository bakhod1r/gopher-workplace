package errorlist

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	t.Run("empty_message", func(t *testing.T) {
		if got := (Errors{}).Error(); got != "" {
			t.Errorf("Error() = %q, want %q", got, "")
		}
	})

	t.Run("single_message", func(t *testing.T) {
		if got := (Errors{ErrA}).Error(); got != "rule a" {
			t.Errorf("Error() = %q, want %q", got, "rule a")
		}
	})

	t.Run("joined_message", func(t *testing.T) {
		if got := (Errors{ErrA, ErrB}).Error(); got != "rule a; rule b" {
			t.Errorf("Error() = %q, want %q", got, "rule a; rule b")
		}
	})

	t.Run("matches_members", func(t *testing.T) {
		var err error = Errors{ErrA, ErrB}
		if !errors.Is(err, ErrA) || !errors.Is(err, ErrB) {
			t.Errorf("errors.Is failed for %v", err)
		}
	})

	t.Run("does_not_match_others", func(t *testing.T) {
		var err error = Errors{ErrA}
		if errors.Is(err, ErrB) {
			t.Error("errors.Is matched a non-member")
		}
	})
}
