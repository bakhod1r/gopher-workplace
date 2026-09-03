package temperr

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsTemporary(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"temporary", &NetError{Temp: true}, true},
		{"permanent", &NetError{Temp: false}, false},
		{"wrapped_temporary", fmt.Errorf("dial: %w", &NetError{Temp: true}), true},
		{"wrapped_permanent", fmt.Errorf("dial: %w", &NetError{Temp: false}), false},
		{"plain", errors.New("boom"), false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsTemporary(tc.err); got != tc.want {
				t.Errorf("IsTemporary(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
