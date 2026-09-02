package requiretext

import (
	"errors"
	"testing"
)

func TestRequire(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want error
	}{
		{"text", "hello", nil},
		{"padded_text", "  hi  ", nil},
		{"empty", "", ErrRequired},
		{"spaces", "   ", ErrRequired},
		{"tab_newline", "\t\n", ErrRequired},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Require(tc.s); !errors.Is(got, tc.want) {
				t.Errorf("Require(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
