package errmsg

import (
	"errors"
	"testing"
)

func TestMessage(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, ""},
		{"sentinel", ErrTimeout, "timeout"},
		{"fresh", errors.New("disk full"), "disk full"},
		{"empty_message", errors.New(""), ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Message(tc.err); got != tc.want {
				t.Errorf("Message(%v) = %q, want %q", tc.err, got, tc.want)
			}
		})
	}
}
