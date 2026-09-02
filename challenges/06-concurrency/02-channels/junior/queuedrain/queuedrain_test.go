package queuedrain

import "testing"

func chanOf(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestDrainQueue(t *testing.T) {
	cases := []struct {
		name     string
		attempts []int
		want     int
	}{
		{"three_pending", []int{1, 2, 3}, 3},
		{"empty_queue", nil, 0},
		{"one_pending", []int{9}, 1},
		{"four_zero_ids", []int{0, 0, 0, 0}, 4},
		{"two_pending", []int{-1, -2}, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := DrainQueue(chanOf(tc.attempts...)); got != tc.want {
				t.Errorf("DrainQueue(%v) = %d, want %d", tc.attempts, got, tc.want)
			}
		})
	}
}
