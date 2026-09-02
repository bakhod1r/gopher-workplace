package percent

import (
	"errors"
	"testing"
)

func TestPercent(t *testing.T) {
	cases := []struct {
		name        string
		part, total int
		want        float64
		wantErr     error
	}{
		{"quarter", 1, 4, 25, nil},
		{"half", 5, 10, 50, nil},
		{"full", 4, 4, 100, nil},
		{"none", 0, 4, 0, nil},
		{"zero_total", 1, 0, 0, ErrZeroTotal},
		{"negative_part", -1, 4, 0, ErrNegative},
		{"negative_total", 1, -4, 0, ErrNegative},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Percent(tc.part, tc.total)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Percent(%d, %d) err = %v, want %v", tc.part, tc.total, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Percent(%d, %d) = %v, want %v", tc.part, tc.total, got, tc.want)
			}
		})
	}
}
