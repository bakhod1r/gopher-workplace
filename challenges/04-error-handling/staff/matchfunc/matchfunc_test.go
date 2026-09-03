package matchfunc

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestAny(t *testing.T) {
	isTimeout := func(e error) bool { return strings.Contains(e.Error(), "timeout") }

	cases := []struct {
		name string
		err  error
		pred func(error) bool
		want bool
	}{
		{"nil", nil, isTimeout, false},
		{"leaf_match", ErrB, isTimeout, true},
		{"leaf_no_match", ErrA, isTimeout, false},
		{"wrapped", fmt.Errorf("call: %w", ErrB), isTimeout, true},
		{"joined", errors.Join(ErrA, ErrB), isTimeout, true},
		{"nested", errors.Join(ErrA, fmt.Errorf("x: %w", errors.Join(ErrB))), isTimeout, true},
		{"none", errors.Join(ErrA, errors.New("boom")), isTimeout, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Any(tc.err, tc.pred); got != tc.want {
				t.Errorf("Any(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
