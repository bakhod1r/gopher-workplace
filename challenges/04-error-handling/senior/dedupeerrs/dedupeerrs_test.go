package dedupeerrs

import (
	"errors"
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	sameMsg := errors.New("shard unreachable")

	t.Run("nil", func(t *testing.T) {
		if got := Unique(nil); got != nil {
			t.Errorf("Unique(nil) = %v, want nil", got)
		}
	})

	t.Run("all_nil", func(t *testing.T) {
		if got := Unique([]error{nil, nil}); got != nil {
			t.Errorf("Unique = %v, want nil", got)
		}
	})

	t.Run("dedupes_by_message", func(t *testing.T) {
		got := Unique([]error{ErrA, sameMsg, ErrB})
		want := []error{ErrA, ErrB}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique = %v, want %v", got, want)
		}
	})

	t.Run("keeps_first_occurrence", func(t *testing.T) {
		got := Unique([]error{ErrB, ErrA, ErrB})
		want := []error{ErrB, ErrA}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique = %v, want %v", got, want)
		}
	})
}
