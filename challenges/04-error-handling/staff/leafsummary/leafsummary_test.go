package leafsummary

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSummary(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want map[string]int
	}{
		{"nil", nil, map[string]int{}},
		{"leaf", ErrA, map[string]int{"a": 1}},
		{"wrapped_counts_leaf_only", fmt.Errorf("x: %w", ErrA), map[string]int{"a": 1}},
		{"joined", errors.Join(ErrA, ErrA, ErrB), map[string]int{"a": 2, "b": 1}},
		{
			"nested",
			errors.Join(fmt.Errorf("x: %w", ErrA), errors.Join(ErrB, ErrB)),
			map[string]int{"a": 1, "b": 2},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Summary(tc.err)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Summary(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
