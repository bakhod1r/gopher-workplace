package emailqueue

import "testing"

func TestSendCampaign(t *testing.T) {
	cost := func(addr string) int { return len(addr) }

	cases := []struct {
		name       string
		recipients []string
		limit      int
		want       int
	}{
		{"two_limit_two", []string{"a", "bb"}, 2, 3},
		{"single_limit_one", []string{"abcd"}, 1, 4},
		{"limit_above_len", []string{"ab", "cd"}, 10, 4},
		{"many_recipients", []string{"a", "b", "c", "d", "e"}, 2, 5},
		{"empty_campaign", nil, 3, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SendCampaign(tc.recipients, tc.limit, cost); got != tc.want {
				t.Errorf("SendCampaign(%v, %d) = %d, want %d", tc.recipients, tc.limit, got, tc.want)
			}
		})
	}
}
