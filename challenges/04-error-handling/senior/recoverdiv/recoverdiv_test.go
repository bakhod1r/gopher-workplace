package recoverdiv

import (
	"errors"
	"testing"
)

func TestSafeDivide(t *testing.T) {
	cases := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{"exact", 10, 2, 5, nil},
		{"negative", -8, 4, -2, nil},
		{"zero_numerator", 0, 5, 0, nil},
		{"zero_divisor", 1, 0, 0, ErrPanic},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SafeDivide(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("SafeDivide(%d, %d) err = %v, want %v", tc.a, tc.b, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("SafeDivide(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
