package asinterface

import (
	"errors"
	"fmt"
	"testing"
)

func TestDelayOf(t *testing.T) {
	cases := []struct {
		name      string
		err       error
		wantDelay int
		wantOK    bool
	}{
		{"nil", nil, 0, false},
		{"direct", &Throttled{Seconds: 5}, 5, true},
		{"wrapped", fmt.Errorf("call: %w", &Throttled{Seconds: 2}), 2, true},
		{"joined", errors.Join(ErrOther, &Throttled{Seconds: 3}), 3, true},
		{"absent", ErrOther, 0, false},
		{"joined_absent", errors.Join(ErrOther, errors.New("x")), 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d, ok := DelayOf(tc.err)
			if ok != tc.wantOK || d != tc.wantDelay {
				t.Errorf("DelayOf(%v) = %d, %v, want %d, %v", tc.err, d, ok, tc.wantDelay, tc.wantOK)
			}
		})
	}
}
