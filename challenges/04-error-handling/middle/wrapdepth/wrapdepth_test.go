package wrapdepth

import (
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
		{"leaf", ErrBase, 1},
		{"one_wrap", fmt.Errorf("a: %w", ErrBase), 2},
		{"two_wraps", fmt.Errorf("a: %w", fmt.Errorf("b: %w", ErrBase)), 3},
		{"plain_message", fmt.Errorf("no verb here"), 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Depth(tc.err); got != tc.want {
				t.Errorf("Depth(%v) = %d, want %d", tc.err, got, tc.want)
			}
		})
	}
}
