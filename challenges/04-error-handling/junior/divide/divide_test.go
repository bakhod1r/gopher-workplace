package divide

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{"exact", 10, 2, 5, nil},
		{"truncated", 7, 2, 3, nil},
		{"zero_divisor", 1, 0, 0, ErrDivideByZero},
		{"negative", -9, 3, -3, nil},
		{"zero_numerator", 0, 5, 0, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Divide(%d, %d) err = %v, want %v", tc.a, tc.b, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Divide(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
