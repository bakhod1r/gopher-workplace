package joinsplit

import (
	"errors"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want []error
	}{
		{"nil", nil, nil},
		{"single", ErrA, []error{ErrA}},
		{"joined_two", errors.Join(ErrA, ErrB), []error{ErrA, ErrB}},
		{"joined_one", errors.Join(ErrA), []error{ErrA}},
		{"joined_with_nil", errors.Join(nil, ErrB), []error{ErrB}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Split(tc.err)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Split(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
