package customas

import (
	"errors"
	"fmt"
	"testing"
)

func TestLegacyAs(t *testing.T) {
	t.Run("converts", func(t *testing.T) {
		var m *Modern
		if !errors.As(&LegacyError{Op: "read", Num: 5}, &m) {
			t.Fatal("errors.As = false, want true")
		}
		if m.Op != "read" || m.Code != 5 {
			t.Errorf("m = %+v, want {read 5}", *m)
		}
	})

	t.Run("through_wrapping", func(t *testing.T) {
		var m *Modern
		err := fmt.Errorf("layer: %w", &LegacyError{Op: "write", Num: 9})
		if !errors.As(err, &m) {
			t.Fatal("errors.As = false, want true")
		}
		if m.Code != 9 {
			t.Errorf("Code = %d, want 9", m.Code)
		}
	})

	t.Run("other_target", func(t *testing.T) {
		var le *LegacyError
		if !errors.As(&LegacyError{Op: "read", Num: 5}, &le) {
			t.Error("errors.As for the concrete type = false, want true")
		}
	})

	t.Run("unrelated_target", func(t *testing.T) {
		type Unrelated struct{ error }
		var u *Unrelated
		if errors.As(&LegacyError{Op: "read"}, &u) {
			t.Error("errors.As matched an unrelated target")
		}
	})
}
