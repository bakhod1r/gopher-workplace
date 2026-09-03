package nilsafejoin

import (
	"errors"
	"testing"
)

func TestClean(t *testing.T) {
	t.Run("nothing", func(t *testing.T) {
		if got := Clean(); got != nil {
			t.Errorf("Clean() = %v, want nil", got)
		}
	})

	t.Run("untyped_nils", func(t *testing.T) {
		if got := Clean(nil, nil); got != nil {
			t.Errorf("Clean = %v, want nil", got)
		}
	})

	t.Run("typed_nil", func(t *testing.T) {
		var typed *OpError
		if got := Clean(typed); got != nil {
			t.Errorf("Clean = %v (%T), want nil", got, got)
		}
	})

	t.Run("mixed", func(t *testing.T) {
		var typed *OpError
		real := &OpError{Op: "read"}
		got := Clean(nil, typed, real)
		if got == nil {
			t.Fatal("got nil, want the real failure")
		}
		if !errors.Is(got, real) {
			t.Errorf("got = %v, want it to match the real failure", got)
		}
		if got.Error() != "read failed" {
			t.Errorf("message = %q, want %q", got.Error(), "read failed")
		}
	})
}
