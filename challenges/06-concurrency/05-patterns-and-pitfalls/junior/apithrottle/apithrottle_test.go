package apithrottle

import "testing"

func TestPeakInFlight(t *testing.T) {
	do := func(string) {}

	cases := []struct {
		name     string
		requests []string
		limit    int
		low      int
		high     int
	}{
		{"five_limit_two", []string{"a", "b", "c", "d", "e"}, 2, 1, 2},
		{"three_limit_one", []string{"a", "b", "c"}, 1, 1, 1},
		{"single_request", []string{"a"}, 5, 1, 1},
		{"three_limit_three", []string{"a", "b", "c"}, 3, 1, 3},
		{"no_requests", nil, 4, 0, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := PeakInFlight(tc.requests, tc.limit, do)
			if got < tc.low || got > tc.high {
				t.Errorf("PeakInFlight(%v, %d) = %d, want between %d and %d",
					tc.requests, tc.limit, got, tc.low, tc.high)
			}
		})
	}
}
