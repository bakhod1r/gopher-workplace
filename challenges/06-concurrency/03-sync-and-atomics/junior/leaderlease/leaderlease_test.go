package leaderlease

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestClaim(t *testing.T) {
	cases := []struct {
		name  string
		steps []int64
		want  []bool
	}{
		{"first_term", []int64{1}, []bool{true}},
		{"replay_same_term", []int64{1, 1}, []bool{true, false}},
		{"skip_ahead", []int64{5}, []bool{false}},
		{"sequence", []int64{1, 2, 3}, []bool{true, true, true}},
		{"stale_then_next", []int64{1, 1, 2}, []bool{true, false, true}},
		{"zero_is_stale", []int64{0}, []bool{false}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var l Lease
			for i, term := range tc.steps {
				if got := l.Claim(term); got != tc.want[i] {
					t.Fatalf("step %d: Claim(%d) = %v, want %v", i, term, got, tc.want[i])
				}
			}
		})
	}
}

func TestTermAdvances(t *testing.T) {
	var l Lease
	if got := l.Term(); got != 0 {
		t.Fatalf("fresh Term() = %d, want 0", got)
	}
	l.Claim(1)
	l.Claim(2)
	if got := l.Term(); got != 2 {
		t.Errorf("Term() = %d, want 2", got)
	}
}

func TestOnlyOneReplicaWinsATerm(t *testing.T) {
	var l Lease
	const replicas = 64

	var wins atomic.Int64
	var start, wg sync.WaitGroup
	start.Add(1)
	for range replicas {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start.Wait()
			if l.Claim(1) {
				wins.Add(1)
			}
		}()
	}
	start.Done()
	wg.Wait()

	if got := wins.Load(); got != 1 {
		t.Errorf("winners = %d, want exactly 1", got)
	}
}
