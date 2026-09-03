package sanitize

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestPublic(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want error
	}{
		{"nil", nil, nil},
		{"missing", errInternalMissing, ErrPublicNotFound},
		{"wrapped_missing", fmt.Errorf("query: %w", errInternalMissing), ErrPublicNotFound},
		{"parse", errInternalParse, ErrPublicInvalid},
		{"unknown", errors.New("/etc/secret: permission denied"), ErrPublicInternal},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Public(tc.err)
			if got != tc.want {
				t.Fatalf("Public(%v) = %v, want %v", tc.err, got, tc.want)
			}
			if got != nil && strings.Contains(got.Error(), "users") {
				t.Errorf("public message %q leaked an internal detail", got.Error())
			}
			if got != nil && strings.Contains(got.Error(), "/etc") {
				t.Errorf("public message %q leaked a path", got.Error())
			}
		})
	}
}
