package customis

import (
	"errors"
	"fmt"
	"testing"
)

func TestStatusError(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		if got := (&StatusError{Code: 503}).Error(); got != "status 503" {
			t.Errorf("Error() = %q, want %q", got, "status 503")
		}
	})

	cases := []struct {
		name   string
		err    error
		target error
		want   bool
	}{
		{"exact", &StatusError{503}, &StatusError{503}, true},
		{"class_marker", &StatusError{503}, &StatusError{500}, true},
		{"other_class", &StatusError{404}, &StatusError{500}, false},
		{"class_404", &StatusError{404}, &StatusError{400}, true},
		{"non_status_target", &StatusError{503}, errors.New("boom"), false},
		{"wrapped", fmt.Errorf("call: %w", &StatusError{503}), &StatusError{500}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := errors.Is(tc.err, tc.target); got != tc.want {
				t.Errorf("errors.Is(%v, %v) = %v, want %v", tc.err, tc.target, got, tc.want)
			}
		})
	}
}
