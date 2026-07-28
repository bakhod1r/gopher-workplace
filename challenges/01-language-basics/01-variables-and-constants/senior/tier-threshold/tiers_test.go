package tiers

import "testing"

func TestThresholds(t *testing.T) {
	if Bronze != 100 || Silver != 200 || Gold != 300 {
		t.Fatalf("tiers=%d,%d,%d; want 100,200,300", Bronze, Silver, Gold)
	}
}
