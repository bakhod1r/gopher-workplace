package sqrtneg

import (
	"errors"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := []struct {
		name    string
		x       float64
		want    float64
		wantErr error
	}{
		{"nine", 9, 3, nil},
		{"zero", 0, 0, nil},
		{"quarter", 0.25, 0.5, nil},
		{"negative", -1, 0, ErrNegative},
		{"big_negative", -100, 0, ErrNegative},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Sqrt(tc.x)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Sqrt(%v) err = %v, want %v", tc.x, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Sqrt(%v) = %v, want %v", tc.x, got, tc.want)
			}
		})
	}
}
