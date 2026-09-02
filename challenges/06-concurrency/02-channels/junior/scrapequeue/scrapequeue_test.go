package scrapequeue

import "testing"

func TestQueueStats(t *testing.T) {
	cases := []struct {
		name    string
		size    int
		wantLen int
		wantCap int
	}{
		{"three", 3, 3, 3},
		{"zero", 0, 0, 0},
		{"negative", -1, 0, 0},
		{"one", 1, 1, 1},
		{"ten", 10, 10, 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotLen, gotCap := QueueStats(tc.size)
			if gotLen != tc.wantLen || gotCap != tc.wantCap {
				t.Errorf("QueueStats(%d) = %d, %d, want %d, %d",
					tc.size, gotLen, gotCap, tc.wantLen, tc.wantCap)
			}
		})
	}
}
