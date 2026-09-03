package isany

import (
	"fmt"
	"testing"
)

func TestIsAny(t *testing.T) {
	cases := []struct {
		name    string
		err     error
		targets []error
		want    bool
	}{
		{"first_target", ErrTimeout, []error{ErrTimeout, ErrReset}, true},
		{"second_target", ErrReset, []error{ErrTimeout, ErrReset}, true},
		{"wrapped", fmt.Errorf("dial: %w", ErrReset), []error{ErrTimeout, ErrReset}, true},
		{"no_match", ErrFatal, []error{ErrTimeout, ErrReset}, false},
		{"no_targets", ErrTimeout, nil, false},
		{"nil_error", nil, []error{ErrTimeout}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsAny(tc.err, tc.targets...); got != tc.want {
				t.Errorf("IsAny(%v, %v) = %v, want %v", tc.err, tc.targets, got, tc.want)
			}
		})
	}
}
