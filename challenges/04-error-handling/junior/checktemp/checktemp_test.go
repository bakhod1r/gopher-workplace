package checktemp

import (
	"errors"
	"testing"
)

func TestCheckTemp(t *testing.T) {
	cases := []struct {
		name string
		c    float64
		want error
	}{
		{"room", 20, nil},
		{"lower_bound", -40, nil},
		{"upper_bound", 85, nil},
		{"below", -50, ErrBelowRange},
		{"just_below", -40.1, ErrBelowRange},
		{"above", 100, ErrAboveRange},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CheckTemp(tc.c); !errors.Is(got, tc.want) {
				t.Errorf("CheckTemp(%v) = %v, want %v", tc.c, got, tc.want)
			}
		})
	}
}
