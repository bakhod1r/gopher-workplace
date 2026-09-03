package mustatoipanic

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func TestMustParse(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		if got := MustParse("42"); got != 42 {
			t.Errorf("MustParse(\"42\") = %d, want 42", got)
		}
	})

	t.Run("negative", func(t *testing.T) {
		if got := MustParse("-7"); got != -7 {
			t.Errorf("MustParse(\"-7\") = %d, want -7", got)
		}
	})

	t.Run("panics_with_wrapped_error", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("MustParse did not panic")
			}
			err, ok := r.(error)
			if !ok {
				t.Fatalf("panic value %v (%T), want an error", r, r)
			}
			if !errors.Is(err, strconv.ErrSyntax) {
				t.Errorf("errors.Is(%v, strconv.ErrSyntax) = false, want true", err)
			}
			if !strings.Contains(err.Error(), `"x"`) {
				t.Errorf("message = %q, want it to quote the input", err.Error())
			}
		}()
		MustParse("x")
	})
}
