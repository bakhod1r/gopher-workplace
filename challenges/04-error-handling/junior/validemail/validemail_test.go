package validemail

import (
	"errors"
	"testing"
)

func TestValidEmail(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want error
	}{
		{"valid", "a@b.com", nil},
		{"no_at", "abc", ErrNoAt},
		{"empty", "", ErrNoAt},
		{"empty_local", "@b.com", ErrEmptyPart},
		{"empty_domain", "a@", ErrEmptyPart},
		{"only_at", "@", ErrEmptyPart},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ValidEmail(tc.s); !errors.Is(got, tc.want) {
				t.Errorf("ValidEmail(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
