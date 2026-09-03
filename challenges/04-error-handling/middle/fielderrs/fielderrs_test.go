package fielderrs

import (
	"errors"
	"testing"
)

func TestCombine(t *testing.T) {
	t.Run("nil_map", func(t *testing.T) {
		if got := Combine(nil); got != nil {
			t.Errorf("Combine(nil) = %v, want nil", got)
		}
	})

	t.Run("all_nil", func(t *testing.T) {
		if got := Combine(map[string]error{"a": nil}); got != nil {
			t.Errorf("Combine = %v, want nil", got)
		}
	})

	t.Run("sorted_message", func(t *testing.T) {
		got := Combine(map[string]error{"b": ErrB, "a": ErrA})
		if got == nil {
			t.Fatal("got nil, want an error")
		}
		want := "a: bad a\nb: bad b"
		if got.Error() != want {
			t.Errorf("message = %q, want %q", got.Error(), want)
		}
	})

	t.Run("matchable", func(t *testing.T) {
		got := Combine(map[string]error{"a": ErrA, "b": ErrB})
		if !errors.Is(got, ErrA) || !errors.Is(got, ErrB) {
			t.Errorf("errors.Is failed for %v", got)
		}
	})

	t.Run("single_field", func(t *testing.T) {
		got := Combine(map[string]error{"a": ErrA, "b": nil})
		if got.Error() != "a: bad a" {
			t.Errorf("message = %q, want %q", got.Error(), "a: bad a")
		}
	})
}
