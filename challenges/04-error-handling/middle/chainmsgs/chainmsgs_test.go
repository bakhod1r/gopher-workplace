package chainmsgs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMessages(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want []string
	}{
		{"nil", nil, nil},
		{"leaf", ErrBase, []string{"base failure"}},
		{"one_wrap", fmt.Errorf("a: %w", ErrBase), []string{"a: base failure", "base failure"}},
		{
			"two_wraps",
			fmt.Errorf("b: %w", fmt.Errorf("a: %w", ErrBase)),
			[]string{"b: a: base failure", "a: base failure", "base failure"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Messages(tc.err)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Messages(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
