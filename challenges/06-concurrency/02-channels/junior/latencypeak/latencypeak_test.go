package latencypeak

import "testing"

func chanOf(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestPeakLatency(t *testing.T) {
	cases := []struct {
		name    string
		samples []int
		wantV   int
		wantOK  bool
	}{
		{"three_samples", []int{30, 90, 40}, 90, true},
		{"empty_window", nil, 0, false},
		{"all_negative", []int{-5, -2}, -2, true},
		{"single", []int{42}, 42, true},
		{"peak_first", []int{10, 1, 2}, 10, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotV, gotOK := PeakLatency(chanOf(tc.samples...))
			if gotV != tc.wantV || gotOK != tc.wantOK {
				t.Errorf("PeakLatency(%v) = %d, %t, want %d, %t",
					tc.samples, gotV, gotOK, tc.wantV, tc.wantOK)
			}
		})
	}
}
