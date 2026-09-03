package joinorder

import (
	"testing"
)

func TestSorted(t *testing.T) {
	t.Run("nothing", func(t *testing.T) {
		if got := Sorted(); got != nil {
			t.Errorf("Sorted() = %v, want nil", got)
		}
	})

	t.Run("all_nil", func(t *testing.T) {
		if got := Sorted(nil, nil); got != nil {
			t.Errorf("Sorted = %v, want nil", got)
		}
	})

	t.Run("orders_by_message", func(t *testing.T) {
		got := Sorted(ErrC, ErrA, ErrB)
		want := "a\nb\nc"
		if got.Error() != want {
			t.Errorf("message = %q, want %q", got.Error(), want)
		}
	})

	t.Run("keeps_duplicates", func(t *testing.T) {
		got := Sorted(ErrA, ErrA)
		if got.Error() != "a\na" {
			t.Errorf("message = %q, want %q", got.Error(), "a\na")
		}
	})

	t.Run("skips_nil", func(t *testing.T) {
		got := Sorted(ErrB, nil, ErrA)
		if got.Error() != "a\nb" {
			t.Errorf("message = %q, want %q", got.Error(), "a\nb")
		}
	})
}
