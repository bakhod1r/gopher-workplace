package mapwrap

import (
	"errors"
	"testing"
)

func TestFind(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}

	t.Run("present", func(t *testing.T) {
		got, err := Find(m, "a")
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if got != 1 {
			t.Errorf("got = %d, want 1", got)
		}
	})

	t.Run("stored_zero", func(t *testing.T) {
		got, err := Find(m, "zero")
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if got != 0 {
			t.Errorf("got = %d, want 0", got)
		}
	})

	t.Run("absent_message", func(t *testing.T) {
		_, err := Find(m, "missing")
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		if err.Error() != "key missing: not found" {
			t.Errorf("message = %q, want %q", err.Error(), "key missing: not found")
		}
	})

	t.Run("absent_matchable", func(t *testing.T) {
		got, err := Find(nil, "a")
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("errors.Is = false for %v", err)
		}
		if got != 0 {
			t.Errorf("got = %d, want 0", got)
		}
	})
}
