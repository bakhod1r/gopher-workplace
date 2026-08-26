package fmtstringer

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	cases := []struct {
		name string
		p    Point
		want string
	}{
		{"positive", Point{1, 2}, "(1,2)"},
		{"origin", Point{0, 0}, "(0,0)"},
		{"negative", Point{-5, -10}, "(-5,-10)"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Test the method directly.
			if got := tc.p.String(); got != tc.want {
				t.Errorf("String() = %q, want %q", got, tc.want)
			}
			// Test that fmt uses it via %v and %s.
			if got := fmt.Sprintf("%v", tc.p); got != tc.want {
				t.Errorf("fmt.Sprintf(%%v) = %q, want %q", got, tc.want)
			}
		})
	}
}
