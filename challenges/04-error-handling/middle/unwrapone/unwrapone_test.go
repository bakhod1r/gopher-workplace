package unwrapone

import (
	"fmt"
	"testing"
)

func TestCause(t *testing.T) {
	inner := fmt.Errorf("inner: %w", ErrBase)

	cases := []struct {
		name string
		err  error
		want error
	}{
		{"nil", nil, nil},
		{"unwrapped", ErrBase, nil},
		{"one_layer", fmt.Errorf("a: %w", ErrBase), ErrBase},
		{"two_layers", fmt.Errorf("b: %w", inner), inner},
		{"no_verb", fmt.Errorf("plain message"), nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Cause(tc.err); got != tc.want {
				t.Errorf("Cause(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
