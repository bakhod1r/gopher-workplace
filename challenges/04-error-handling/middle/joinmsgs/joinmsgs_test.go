package joinmsgs

import (
	"errors"
	"reflect"
	"testing"
)

func TestLines(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want []string
	}{
		{"nil", nil, nil},
		{"single", ErrA, []string{"a"}},
		{"joined", errors.Join(ErrA, ErrB), []string{"a", "b"}},
		{"joined_one", errors.Join(ErrA), []string{"a"}},
		{"joined_with_nil", errors.Join(nil, ErrB), []string{"b"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Lines(tc.err)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Lines(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
