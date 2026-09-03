package chainto

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAbove(t *testing.T) {
	t.Run("target_is_the_error", func(t *testing.T) {
		got := Above(ErrA, ErrA)
		if got == nil {
			t.Fatal("got nil, want an empty slice")
		}
		if len(got) != 0 {
			t.Errorf("got = %v, want empty", got)
		}
	})

	t.Run("one_layer", func(t *testing.T) {
		got := Above(fmt.Errorf("x: %w", ErrA), ErrA)
		want := []string{"x: a"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got = %v, want %v", got, want)
		}
	})

	t.Run("two_layers", func(t *testing.T) {
		got := Above(fmt.Errorf("y: %w", fmt.Errorf("x: %w", ErrA)), ErrA)
		want := []string{"y: x: a", "x: a"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got = %v, want %v", got, want)
		}
	})

	t.Run("absent", func(t *testing.T) {
		if got := Above(fmt.Errorf("x: %w", ErrB), ErrA); got != nil {
			t.Errorf("got = %v, want nil", got)
		}
	})

	t.Run("nil_error", func(t *testing.T) {
		if got := Above(nil, ErrA); got != nil {
			t.Errorf("got = %v, want nil", got)
		}
	})
}
