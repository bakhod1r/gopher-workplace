package validage

import (
	"errors"
	"testing"
)

func TestValidAge(t *testing.T) {
	cases := []struct {
		name string
		age  int
		want error
	}{
		{"typical", 30, nil},
		{"zero", 0, nil},
		{"upper_bound", 130, nil},
		{"negative", -1, ErrTooYoung},
		{"too_old", 200, ErrTooOld},
		{"just_over", 131, ErrTooOld},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ValidAge(tc.age); !errors.Is(got, tc.want) {
				t.Errorf("ValidAge(%d) = %v, want %v", tc.age, got, tc.want)
			}
		})
	}
}
