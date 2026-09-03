package treedepth

import (
	"errors"
	"fmt"
	"testing"
)

func TestDepth(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want int
	}{
		{"nil", nil, 0},
		{"leaf", ErrA, 1},
		{"one_wrap", fmt.Errorf("x: %w", ErrA), 2},
		{"two_wraps", fmt.Errorf("y: %w", fmt.Errorf("x: %w", ErrA)), 3},
		{"join_of_leaves", errors.Join(ErrA, ErrB), 2},
		{"join_with_wrap", errors.Join(ErrA, fmt.Errorf("x: %w", ErrB)), 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Depth(tc.err); got != tc.want {
				t.Errorf("Depth(%v) = %d, want %d", tc.err, got, tc.want)
			}
		})
	}
}
