package licenseseats

import (
	"sync"
	"testing"
)

func TestSeatPool(t *testing.T) {
	cases := []struct {
		name     string
		seats    int64
		acquires int
		releases int
		wantLast bool
		wantFree int64
	}{
		{"one_seat", 1, 1, 0, true, 0},
		{"exhausted", 1, 2, 0, false, 0},
		{"no_seats", 0, 1, 0, false, 0},
		{"released", 1, 1, 1, true, 1},
		{"partial_use", 3, 2, 0, true, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewSeatPool(tc.seats)
			var last bool
			for i := 0; i < tc.acquires; i++ {
				last = p.TryAcquire()
			}
			for i := 0; i < tc.releases; i++ {
				p.Release()
			}
			if tc.acquires > 0 && last != tc.wantLast {
				t.Errorf("last TryAcquire() = %v, want %v", last, tc.wantLast)
			}
			if got := p.Free(); got != tc.wantFree {
				t.Errorf("Free() = %d, want %d", got, tc.wantFree)
			}
		})
	}
}

func TestSeatPoolConcurrent(t *testing.T) {
	const seats = 8
	const logins = 40
	p := NewSeatPool(seats)

	granted := make(chan bool, logins)
	var wg sync.WaitGroup
	wg.Add(logins)
	for i := 0; i < logins; i++ {
		go func() {
			defer wg.Done()
			granted <- p.TryAcquire()
			p.Free()
		}()
	}
	wg.Wait()
	close(granted)

	n := 0
	for g := range granted {
		if g {
			n++
		}
	}
	if n != seats {
		t.Errorf("seats granted = %d, want %d", n, seats)
	}
	if got := p.Free(); got != 0 {
		t.Errorf("Free() = %d, want 0", got)
	}
}
