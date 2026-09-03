package opaqueerr

import (
	"errors"
	"fmt"
	"testing"
)

func TestPredicates(t *testing.T) {
	cases := []struct {
		name          string
		err           error
		wantRetryable bool
		wantRejected  bool
	}{
		{"nil", nil, false, false},
		{"transient", ErrTransient, true, false},
		{"invalid", ErrInvalid, false, true},
		{"wrapped_transient", fmt.Errorf("call: %w", ErrTransient), true, false},
		{"wrapped_invalid", fmt.Errorf("call: %w", ErrInvalid), false, true},
		{"unknown", errors.New("boom"), false, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsRetryable(tc.err); got != tc.wantRetryable {
				t.Errorf("IsRetryable(%v) = %v, want %v", tc.err, got, tc.wantRetryable)
			}
			if got := IsRejected(tc.err); got != tc.wantRejected {
				t.Errorf("IsRejected(%v) = %v, want %v", tc.err, got, tc.wantRejected)
			}
		})
	}
}
