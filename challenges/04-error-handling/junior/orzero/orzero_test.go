package orzero

import (
	"errors"
	"testing"
)

func TestOrZero(t *testing.T) {
	cases := []struct {
		name string
		v    int
		err  error
		want int
	}{
		{"ok", 42, nil, 42},
		{"ok_zero", 0, nil, 0},
		{"failed", 42, ErrHost, 0},
		{"failed_negative", -7, errors.New("boom"), 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := OrZero(tc.v, tc.err); got != tc.want {
				t.Errorf("OrZero(%d, %v) = %d, want %d", tc.v, tc.err, got, tc.want)
			}
		})
	}
}
