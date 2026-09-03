package isnotfound

import (
	"fmt"
	"testing"
)

func TestIsNotFound(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"direct", ErrNotFound, true},
		{"wrapped_once", fmt.Errorf("load: %w", ErrNotFound), true},
		{"wrapped_twice", fmt.Errorf("handler: %w", fmt.Errorf("load: %w", ErrNotFound)), true},
		{"other_sentinel", ErrDenied, false},
		{"wrapped_other", fmt.Errorf("load: %w", ErrDenied), false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNotFound(tc.err); got != tc.want {
				t.Errorf("IsNotFound(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
