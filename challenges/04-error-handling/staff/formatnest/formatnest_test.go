package formatnest

import (
	"fmt"
	"testing"
)

func TestVerbose(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, ""},
		{"leaf", ErrA, "a"},
		{"one_wrap", fmt.Errorf("x: %w", ErrA), "x: a\ncaused by: a"},
		{
			"two_wraps",
			fmt.Errorf("y: %w", fmt.Errorf("x: %w", ErrA)),
			"y: x: a\ncaused by: x: a\ncaused by: a",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Verbose(tc.err); got != tc.want {
				t.Errorf("Verbose(%v) = %q, want %q", tc.err, got, tc.want)
			}
		})
	}
}
