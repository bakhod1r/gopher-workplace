package invoicetotal

import "testing"

func TestTotalCents(t *testing.T) {
	cases := []struct {
		name    string
		lines   []int
		workers int
		want    int
	}{
		{"three_lines_two_workers", []int{100, 250, 99}, 2, 449},
		{"empty_invoice", nil, 4, 0},
		{"single_line", []int{500}, 1, 500},
		{"more_workers_than_lines", []int{200, 200}, 8, 400},
		{"zero_workers", []int{1, 2, 3, 4}, 0, 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TotalCents(tc.lines, tc.workers); got != tc.want {
				t.Errorf("TotalCents(%v, %d) = %d, want %d",
					tc.lines, tc.workers, got, tc.want)
			}
		})
	}
}
