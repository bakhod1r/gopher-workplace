package parsepair

import (
	"errors"
	"testing"
)

func TestParsePair(t *testing.T) {
	cases := []struct {
		name    string
		s       string
		wantK   string
		wantV   string
		wantErr error
	}{
		{"basic", "HOST=local", "HOST", "local", nil},
		{"empty_value", "HOST=", "HOST", "", nil},
		{"second_equals", "URL=a=b", "URL", "a=b", nil},
		{"no_separator", "HOST", "", "", ErrNoSeparator},
		{"empty_input", "", "", "", ErrNoSeparator},
		{"empty_key", "=x", "", "", ErrEmptyKey},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			k, v, err := ParsePair(tc.s)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ParsePair(%q) err = %v, want %v", tc.s, err, tc.wantErr)
			}
			if k != tc.wantK || v != tc.wantV {
				t.Errorf("ParsePair(%q) = %q, %q, want %q, %q", tc.s, k, v, tc.wantK, tc.wantV)
			}
		})
	}
}
