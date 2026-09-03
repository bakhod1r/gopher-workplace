package errreport

import (
	"errors"
	"testing"
)

func TestReport(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, ""},
		{"single", ErrA, "1. a"},
		{"joined", errors.Join(ErrA, ErrB), "1. a\n2. b"},
		{"joined_one", errors.Join(ErrA), "1. a"},
		{"joined_with_nil", errors.Join(nil, ErrB), "1. b"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Report(tc.err); got != tc.want {
				t.Errorf("Report(%v) = %q, want %q", tc.err, got, tc.want)
			}
		})
	}
}
