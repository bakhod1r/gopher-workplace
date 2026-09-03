package stripwrap

import (
	"errors"
	"fmt"
	"testing"
)

func TestStrip(t *testing.T) {
	other := errors.New("boom")
	wrappedOther := fmt.Errorf("layer: %w", other)

	cases := []struct {
		name string
		err  error
		want error
	}{
		{"nil", nil, nil},
		{"direct_sentinel", ErrNotFound, ErrNotFound},
		{"wrapped_not_found", fmt.Errorf("a: %w", ErrNotFound), ErrNotFound},
		{"wrapped_denied", fmt.Errorf("a: %w", fmt.Errorf("b: %w", ErrDenied)), ErrDenied},
		{"unknown", other, other},
		{"wrapped_unknown", wrappedOther, wrappedOther},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Strip(tc.err); got != tc.want {
				t.Errorf("Strip(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
