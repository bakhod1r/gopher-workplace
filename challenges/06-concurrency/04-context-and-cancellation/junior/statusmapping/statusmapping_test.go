package statusmapping

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestStatus(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"success", nil, "ok"},
		{"client_hung_up", context.Canceled, "client_closed_request"},
		{"budget_expired", context.DeadlineExceeded, "gateway_timeout"},
		{"wrapped_cancel", fmt.Errorf("upstream: %w", context.Canceled), "client_closed_request"},
		{"wrapped_deadline", fmt.Errorf("query %q: %w", "SELECT 1", context.DeadlineExceeded), "gateway_timeout"},
		{"doubly_wrapped", fmt.Errorf("handler: %w", fmt.Errorf("db: %w", context.DeadlineExceeded)), "gateway_timeout"},
		{"unrelated", errors.New("constraint violation"), "internal_error"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Status(tc.err); got != tc.want {
				t.Errorf("Status(%v) = %q, want %q", tc.err, got, tc.want)
			}
		})
	}
}
