package treestring

import (
	"errors"
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := Tree(nil); got != "" {
			t.Errorf("Tree(nil) = %q, want %q", got, "")
		}
	})

	t.Run("leaf", func(t *testing.T) {
		if got := Tree(ErrA); got != "a" {
			t.Errorf("Tree = %q, want %q", got, "a")
		}
	})

	t.Run("wrapped", func(t *testing.T) {
		want := "x: a\n\ta"
		if got := Tree(fmt.Errorf("x: %w", ErrA)); got != want {
			t.Errorf("Tree = %q, want %q", got, want)
		}
	})

	t.Run("joined", func(t *testing.T) {
		got := Tree(errors.Join(ErrA, ErrB))
		want := "a\nb\n\ta\n\tb"
		if got != want {
			t.Errorf("Tree = %q, want %q", got, want)
		}
	})
}
