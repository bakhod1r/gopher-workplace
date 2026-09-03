package depthlimit

import (
	"errors"
	"fmt"
	"testing"
)

func TestWithin(t *testing.T) {
	one := fmt.Errorf("x: %w", ErrA)
	three := fmt.Errorf("z: %w", fmt.Errorf("y: %w", one))

	cases := []struct {
		name   string
		err    error
		target error
		max    int
		want   bool
	}{
		{"zero_max", ErrA, ErrA, 0, false},
		{"nil_err", nil, ErrA, 3, false},
		{"nil_target", ErrA, nil, 3, false},
		{"direct", ErrA, ErrA, 1, true},
		{"one_link_within_two", one, ErrA, 2, true},
		{"one_link_beyond_one", one, ErrA, 1, false},
		{"four_links_within_four", three, ErrA, 4, true},
		{"four_links_beyond_three", three, ErrA, 3, false},
		{"absent", one, errors.New("other"), 5, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Within(tc.err, tc.target, tc.max); got != tc.want {
				t.Errorf("Within(%v, %v, %d) = %v, want %v", tc.err, tc.target, tc.max, got, tc.want)
			}
		})
	}
}
