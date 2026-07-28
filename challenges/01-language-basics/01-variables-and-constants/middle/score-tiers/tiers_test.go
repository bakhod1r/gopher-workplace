package tiers

import "testing"

func TestThresholds(t *testing.T) {
	if Bronze != 100 || Silver != 200 || Gold != 300 {
		t.Fatalf("tiers = %d,%d,%d; want 100,200,300", Bronze, Silver, Gold)
	}
}

func TestRank(t *testing.T) {
	cases := []struct {
		score int
		want  Tier
	}{
		{50, 0}, {100, Bronze}, {150, Bronze}, {200, Silver}, {350, Gold},
	}
	for _, c := range cases {
		if got := Rank(c.score); got != c.want {
			t.Errorf("Rank(%d)=%d; want %d", c.score, got, c.want)
		}
	}
}
