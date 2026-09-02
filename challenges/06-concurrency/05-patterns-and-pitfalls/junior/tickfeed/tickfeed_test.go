package tickfeed

import "testing"

func TestLiveTicks(t *testing.T) {
	cases := []struct {
		name       string
		ticks      []int
		closeTicks bool
		closeDone  bool
		want       []int
	}{
		{"three_ticks", []int{1, 2, 3}, true, false, []int{1, 2, 3}},
		{"single_tick", []int{42}, true, false, []int{42}},
		{"five_ticks", []int{1, 2, 3, 4, 5}, true, false, []int{1, 2, 3, 4, 5}},
		{"feed_closed_empty", nil, true, false, nil},
		{"consumer_gone_first", nil, false, true, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ticks := make(chan int, len(tc.ticks))
			for _, v := range tc.ticks {
				ticks <- v
			}
			if tc.closeTicks {
				close(ticks)
			}

			done := make(chan struct{})
			if tc.closeDone {
				close(done)
			}

			var got []int
			for v := range LiveTicks(done, ticks) {
				got = append(got, v)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("LiveTicks() = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("LiveTicks() = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
