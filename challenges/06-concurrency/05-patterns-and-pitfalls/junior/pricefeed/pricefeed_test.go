package pricefeed

import "testing"

func TestPriceFeed(t *testing.T) {
	cases := []struct {
		name string
		base int
		take int
		want []int
	}{
		{"from_hundred", 100, 3, []int{100, 101, 102}},
		{"from_zero", 0, 5, []int{0, 1, 2, 3, 4}},
		{"single_quote", 7, 1, []int{7}},
		{"negative_base", -2, 3, []int{-2, -1, 0}},
		{"take_none", 42, 0, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			done := make(chan struct{})
			feed := PriceFeed(done, tc.base)

			var got []int
			for i := 0; i < tc.take; i++ {
				got = append(got, <-feed)
			}
			close(done)

			// The feed must shut down and close its channel, or this drain hangs.
			for range feed {
			}

			if len(got) != len(tc.want) {
				t.Fatalf("PriceFeed(done, %d) first %d = %v, want %v", tc.base, tc.take, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("PriceFeed(done, %d) first %d = %v, want %v", tc.base, tc.take, got, tc.want)
				}
			}
		})
	}
}
