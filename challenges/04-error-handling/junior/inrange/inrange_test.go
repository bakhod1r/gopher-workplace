package inrange

import (
	"errors"
	"testing"
)

func TestInRange(t *testing.T) {
	cases := []struct {
		name      string
		n, lo, hi int
		want      error
	}{
		{"middle", 5, 1, 10, nil},
		{"lower_bound", 1, 1, 10, nil},
		{"upper_bound", 10, 1, 10, nil},
		{"single_point", 3, 3, 3, nil},
		{"below", 0, 1, 10, ErrOutOfRange},
		{"above", 11, 1, 10, ErrOutOfRange},
		{"reversed", 5, 10, 1, ErrBadBounds},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := InRange(tc.n, tc.lo, tc.hi); !errors.Is(got, tc.want) {
				t.Errorf("InRange(%d, %d, %d) = %v, want %v", tc.n, tc.lo, tc.hi, got, tc.want)
			}
		})
	}
}
