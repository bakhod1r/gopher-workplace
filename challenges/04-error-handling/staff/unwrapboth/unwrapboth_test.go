package unwrapboth

import (
	"errors"
	"testing"
)

func TestMulti(t *testing.T) {
	m := &Multi{First: ErrA, Others: []error{ErrB}}

	t.Run("message", func(t *testing.T) {
		if m.Error() != "a" {
			t.Errorf("Error() = %q, want %q", m.Error(), "a")
		}
	})

	t.Run("primary", func(t *testing.T) {
		if m.Primary() != ErrA {
			t.Errorf("Primary() = %v, want ErrA", m.Primary())
		}
	})

	t.Run("matches_all_members", func(t *testing.T) {
		var err error = m
		if !errors.Is(err, ErrA) || !errors.Is(err, ErrB) {
			t.Error("errors.Is failed for a member")
		}
	})

	t.Run("primary_first", func(t *testing.T) {
		members := m.Unwrap()
		if len(members) != 2 || members[0] != ErrA || members[1] != ErrB {
			t.Errorf("Unwrap = %v, want [a b]", members)
		}
	})

	t.Run("no_others", func(t *testing.T) {
		solo := &Multi{First: ErrA}
		if members := solo.Unwrap(); len(members) != 1 || members[0] != ErrA {
			t.Errorf("Unwrap = %v, want [a]", members)
		}
	})
}
