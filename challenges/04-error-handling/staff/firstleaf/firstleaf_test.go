package firstleaf

import (
	"errors"
	"fmt"
	"testing"
)

func TestOrigin(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want error
	}{
		{"nil", nil, nil},
		{"leaf", ErrA, ErrA},
		{"wrapped", fmt.Errorf("x: %w", ErrA), ErrA},
		{"double_wrapped", fmt.Errorf("y: %w", fmt.Errorf("x: %w", ErrA)), ErrA},
		{"joined", errors.Join(ErrA, ErrB), ErrA},
		{"joined_reversed", errors.Join(ErrB, ErrA), ErrB},
		{"nested", errors.Join(fmt.Errorf("x: %w", ErrB), ErrA), ErrB},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Origin(tc.err); got != tc.want {
				t.Errorf("Origin(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
