package formatverb

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	e := &DetailError{Msg: "failed", Detail: "at line 4"}

	t.Run("error_method", func(t *testing.T) {
		if e.Error() != "failed" {
			t.Errorf("Error() = %q, want %q", e.Error(), "failed")
		}
	})

	t.Run("v", func(t *testing.T) {
		if got := fmt.Sprintf("%v", e); got != "failed" {
			t.Errorf("%%v = %q, want %q", got, "failed")
		}
	})

	t.Run("s", func(t *testing.T) {
		if got := fmt.Sprintf("%s", e); got != "failed" {
			t.Errorf("%%s = %q, want %q", got, "failed")
		}
	})

	t.Run("plus_v", func(t *testing.T) {
		want := "failed\n\tat line 4"
		if got := fmt.Sprintf("%+v", e); got != want {
			t.Errorf("%%+v = %q, want %q", got, want)
		}
	})
}
