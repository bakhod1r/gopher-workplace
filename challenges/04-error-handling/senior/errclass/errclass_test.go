package errclass

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestClassify(t *testing.T) {
	cases := []struct {
		name string
		errs []error
		want map[string]int
	}{
		{"nil", nil, map[string]int{}},
		{"all_nil", []error{nil, nil}, map[string]int{}},
		{"timeouts", []error{ErrTimeout, ErrTimeout}, map[string]int{"timeout": 2}},
		{
			"mixed",
			[]error{ErrTimeout, fmt.Errorf("x: %w", ErrDenied), errors.New("boom"), nil},
			map[string]int{"timeout": 1, "denied": 1, "other": 1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Classify(tc.errs)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Classify(%v) = %v, want %v", tc.errs, got, tc.want)
			}
		})
	}
}
