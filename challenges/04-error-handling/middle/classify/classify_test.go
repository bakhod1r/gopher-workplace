package classify

import (
	"errors"
	"fmt"
	"testing"
)

func TestStatus(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want int
	}{
		{"nil", nil, 200},
		{"not_found", ErrNotFound, 404},
		{"denied", ErrDenied, 403},
		{"conflict", ErrConflict, 409},
		{"wrapped_not_found", fmt.Errorf("load: %w", ErrNotFound), 404},
		{"unknown", errors.New("boom"), 500},
		{"wrapped_unknown", fmt.Errorf("load: %w", errors.New("boom")), 500},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Status(tc.err); got != tc.want {
				t.Errorf("Status(%v) = %d, want %d", tc.err, got, tc.want)
			}
		})
	}
}
