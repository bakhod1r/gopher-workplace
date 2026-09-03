package sentineltable

import (
	"errors"
	"fmt"
	"testing"
)

func TestStatus(t *testing.T) {
	table := []Rule{
		{Err: ErrNotFound, Code: 404},
		{Err: ErrDenied, Code: 403},
	}

	cases := []struct {
		name string
		err  error
		want int
	}{
		{"nil", nil, 200},
		{"not_found", ErrNotFound, 404},
		{"denied", ErrDenied, 403},
		{"wrapped", fmt.Errorf("x: %w", ErrDenied), 403},
		{"unknown", errors.New("boom"), 500},
		{"empty_table", ErrNotFound, 500},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tbl := table
			if tc.name == "empty_table" {
				tbl = nil
			}
			if got := Status(tc.err, tbl); got != tc.want {
				t.Errorf("Status(%v) = %d, want %d", tc.err, got, tc.want)
			}
		})
	}
}
