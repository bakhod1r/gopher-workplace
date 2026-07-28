package plan

import "testing"

func TestTiersAreDistinct(t *testing.T) {
	if Free == Pro || Pro == Enterprise || Free == Enterprise {
		t.Fatalf("tiers must be distinct: Free=%d Pro=%d Enterprise=%d",
			Free, Pro, Enterprise)
	}
	if Free != 0 || Pro != 1 || Enterprise != 2 {
		t.Fatalf("tiers must be ascending 0,1,2: Free=%d Pro=%d Enterprise=%d",
			Free, Pro, Enterprise)
	}
}

func TestLimit(t *testing.T) {
	cases := []struct {
		name string
		tier Tier
		want int
	}{
		{"free", Free, 60},
		{"pro", Pro, 600},
		{"enterprise", Enterprise, 6000},
		{"unknown falls back to free", Tier(99), 60},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Limit(tc.tier); got != tc.want {
				t.Errorf("Limit(%v) = %d, want %d", tc.tier, got, tc.want)
			}
		})
	}
}
