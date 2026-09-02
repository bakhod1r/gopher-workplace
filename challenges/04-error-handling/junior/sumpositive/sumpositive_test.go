package sumpositive

import (
	"errors"
	"testing"
)

func TestSumPositive(t *testing.T) {
	cases := []struct {
		name    string
		nums    []int
		want    int
		wantErr error
	}{
		{"nil", nil, 0, nil},
		{"empty", []int{}, 0, nil},
		{"basic", []int{1, 2, 3}, 6, nil},
		{"with_zero", []int{0, 5}, 5, nil},
		{"negative_second", []int{1, -2}, 0, ErrNegativeValue},
		{"negative_last", []int{1, 2, -1}, 0, ErrNegativeValue},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SumPositive(tc.nums)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("SumPositive(%v) err = %v, want %v", tc.nums, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("SumPositive(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
