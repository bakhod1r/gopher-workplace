package requestid

import (
	"sync"
	"testing"
)

func TestIDGenSequential(t *testing.T) {
	cases := []struct {
		name    string
		calls   int
		wantID  int64
		wantAll int64
	}{
		{"no_requests", 0, 0, 0},
		{"first", 1, 1, 1},
		{"second", 2, 2, 2},
		{"tenth", 10, 10, 10},
		{"hundredth", 100, 100, 100},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g IDGen
			var last int64
			for i := 0; i < tc.calls; i++ {
				last = g.Next()
			}
			if last != tc.wantID {
				t.Errorf("last Next() = %d, want %d", last, tc.wantID)
			}
			if got := g.Issued(); got != tc.wantAll {
				t.Errorf("Issued() = %d, want %d", got, tc.wantAll)
			}
		})
	}
}

func TestIDGenUnique(t *testing.T) {
	var g IDGen
	const handlers = 16
	const per = 100
	ids := make(chan int64, handlers*per)
	var wg sync.WaitGroup
	wg.Add(handlers)
	for i := 0; i < handlers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < per; j++ {
				ids <- g.Next()
			}
		}()
	}
	wg.Wait()
	close(ids)

	seen := make(map[int64]bool, handlers*per)
	for id := range ids {
		if seen[id] {
			t.Fatalf("duplicate request ID %d", id)
		}
		seen[id] = true
	}
	if len(seen) != handlers*per {
		t.Errorf("unique IDs = %d, want %d", len(seen), handlers*per)
	}
	if got := g.Issued(); got != int64(handlers*per) {
		t.Errorf("Issued() = %d, want %d", got, handlers*per)
	}
}
