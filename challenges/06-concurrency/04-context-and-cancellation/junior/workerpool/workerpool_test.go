package workerpool

import (
	"context"
	"testing"
)

func TestStopWorkers(t *testing.T) {
	cases := []struct {
		name string
		n    int
	}{
		{"no_workers", 0},
		{"one_worker", 1},
		{"three_workers", 3},
		{"eight_workers", 8},
		{"fifty_workers", 50},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := StopWorkers(tc.n)
			if len(got) != tc.n {
				t.Fatalf("len(StopWorkers(%d)) = %d, want %d", tc.n, len(got), tc.n)
			}
			for i, err := range got {
				if err != context.Canceled {
					t.Errorf("worker %d stopped with %v, want %v", i, err, context.Canceled)
				}
			}
		})
	}
}
