package batchtotal

import "testing"

func TestBatchTotal(t *testing.T) {
	cases := []struct {
		name    string
		amounts []int
		want    int
	}{
		{"three_amounts", []int{100, 250, 99}, 449},
		{"empty_day", nil, 0},
		{"refund_cancels", []int{-200, 200}, 0},
		{"single", []int{9}, 9},
		{"many", []int{1, 1, 1, 1, 1, 1}, 6},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := BatchTotal(tc.amounts); got != tc.want {
				t.Errorf("BatchTotal(%v) = %d, want %d", tc.amounts, got, tc.want)
			}
		})
	}
}
