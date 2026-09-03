package httpcode

import (
	"errors"
	"fmt"
	"testing"
)

func TestCodeOf(t *testing.T) {
	cases := []struct {
		name     string
		err      error
		wantCode int
		wantOK   bool
	}{
		{"nil", nil, 0, false},
		{"direct", &HTTPError{Code: 404}, 404, true},
		{"wrapped", fmt.Errorf("get: %w", &HTTPError{Code: 500}), 500, true},
		{"double_wrapped", fmt.Errorf("a: %w", fmt.Errorf("b: %w", &HTTPError{Code: 503})), 503, true},
		{"other", errors.New("boom"), 0, false},
		{"wrapped_other", fmt.Errorf("get: %w", errors.New("boom")), 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			code, ok := CodeOf(tc.err)
			if ok != tc.wantOK {
				t.Fatalf("CodeOf(%v) ok = %v, want %v", tc.err, ok, tc.wantOK)
			}
			if code != tc.wantCode {
				t.Errorf("CodeOf(%v) code = %d, want %d", tc.err, code, tc.wantCode)
			}
		})
	}
}
