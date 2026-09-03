package nilinterface

import "testing"

func TestWrap(t *testing.T) {
	t.Run("nil_pointer_is_nil_error", func(t *testing.T) {
		if got := Wrap(nil); got != nil {
			t.Errorf("Wrap(nil) = %v (%T), want a nil error", got, got)
		}
	})

	t.Run("typed_nil_variable", func(t *testing.T) {
		var e *OpError
		if got := Wrap(e); got != nil {
			t.Errorf("Wrap(typed nil) = %v (%T), want a nil error", got, got)
		}
	})

	t.Run("non_nil", func(t *testing.T) {
		got := Wrap(&OpError{Op: "read"})
		if got == nil {
			t.Fatal("Wrap returned nil, want an error")
		}
		if got.Error() != "read failed" {
			t.Errorf("message = %q, want %q", got.Error(), "read failed")
		}
	})
}
