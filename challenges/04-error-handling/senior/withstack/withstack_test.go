package withstack

import (
	"errors"
	"regexp"
	"testing"
)

func TestHere(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := Here(nil); got != nil {
			t.Errorf("Here(nil) = %v, want nil", got)
		}
	})

	t.Run("caller_file_and_line", func(t *testing.T) {
		err := Here(ErrBase)
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		re := regexp.MustCompile(`^withstack_test\.go:\d+: base$`)
		if !re.MatchString(err.Error()) {
			t.Errorf("message = %q, want it to match %v", err.Error(), re)
		}
	})

	t.Run("matchable", func(t *testing.T) {
		if !errors.Is(Here(ErrBase), ErrBase) {
			t.Error("errors.Is = false, want true")
		}
	})

	t.Run("distinct_lines", func(t *testing.T) {
		a := Here(ErrBase)
		b := Here(ErrBase)
		if a.Error() == b.Error() {
			t.Errorf("two call sites produced the same message %q", a.Error())
		}
	})
}
