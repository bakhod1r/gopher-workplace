package joinflatten

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestLeaves(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want []error
	}{
		{"nil", nil, nil},
		{"leaf", ErrA, []error{ErrA}},
		{"wrapped", fmt.Errorf("x: %w", ErrA), []error{ErrA}},
		{"joined", errors.Join(ErrA, ErrB), []error{ErrA, ErrB}},
		{"nested_join", errors.Join(ErrA, errors.Join(ErrB, ErrC)), []error{ErrA, ErrB, ErrC}},
		{"join_of_wraps", errors.Join(fmt.Errorf("x: %w", ErrA), ErrB), []error{ErrA, ErrB}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Leaves(tc.err)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Leaves(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
