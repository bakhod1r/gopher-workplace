package rootcause

import (
	"errors"
	"fmt"
	"testing"
)

func TestRoot(t *testing.T) {
	other := errors.New("other")

	cases := []struct {
		name string
		err  error
		want error
	}{
		{"nil", nil, nil},
		{"unwrapped", ErrBase, ErrBase},
		{"one_layer", fmt.Errorf("a: %w", ErrBase), ErrBase},
		{"three_layers", fmt.Errorf("a: %w", fmt.Errorf("b: %w", fmt.Errorf("c: %w", ErrBase))), ErrBase},
		{"other_root", fmt.Errorf("a: %w", other), other},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Root(tc.err); got != tc.want {
				t.Errorf("Root(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
