package erroras

import (
	"errors"
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	if err := Call(200); err != nil {
		t.Errorf("Call(200) = %v, want nil", err)
	}
	err := Call(500)
	if err == nil {
		t.Fatal("Call(500) = nil, want error")
	}
	if got := err.Error(); got != "call: http 500" {
		t.Errorf("Error() = %q, want \"call: http 500\"", got)
	}
}

func TestStatusOf(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want int
	}{
		{"wrapped", Call(404), 404},
		{"double_wrapped", fmt.Errorf("outer: %w", Call(503)), 503},
		{"bare", &HTTPError{Status: 301}, 301},
		{"unrelated", errors.New("nope"), 0},
		{"nil", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := StatusOf(tc.err); got != tc.want {
				t.Errorf("StatusOf = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestRetryable(t *testing.T) {
	if !Retryable(Call(500)) {
		t.Error("500 should be retryable")
	}
	if !Retryable(Call(503)) {
		t.Error("503 should be retryable")
	}
	if Retryable(Call(404)) {
		t.Error("404 should not be retryable")
	}
	if Retryable(nil) {
		t.Error("nil should not be retryable")
	}
}
