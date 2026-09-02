package haserror

import (
	"errors"
	"testing"
)

func TestHasError(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"sentinel", ErrSample, true},
		{"fresh", errors.New("boom"), true},
		{"empty_message", errors.New(""), true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasError(tc.err); got != tc.want {
				t.Errorf("HasError(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
