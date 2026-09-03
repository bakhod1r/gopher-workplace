package warehouseslots

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestStockReserve(t *testing.T) {
	cases := []struct {
		name      string
		units     int64
		reserve   []int64
		wantLast  bool
		wantAvail int64
	}{
		{"reserve_some", 10, []int64{3}, true, 7},
		{"reserve_exactly_all", 5, []int64{5}, true, 0},
		{"reserve_too_many", 2, []int64{5}, false, 2},
		{"sequential_until_empty", 4, []int64{3, 1}, true, 0},
		{"one_past_empty", 4, []int64{4, 1}, false, 0},
		{"zero_units_rejected", 10, []int64{0}, false, 10},
		{"negative_rejected", 10, []int64{-4}, false, 10},
		{"negative_stock_clamped", -3, []int64{1}, false, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStock(tc.units)
			var ok bool
			for _, n := range tc.reserve {
				ok = s.Reserve(n)
			}
			if ok != tc.wantLast {
				t.Errorf("last Reserve = %v, want %v", ok, tc.wantLast)
			}
			if got := s.Available(); got != tc.wantAvail {
				t.Errorf("Available() = %d, want %d", got, tc.wantAvail)
			}
		})
	}
}

func TestStockRelease(t *testing.T) {
	s := NewStock(1)
	if !s.Reserve(1) {
		t.Fatal("Reserve(1) = false, want true")
	}
	s.Release(0)
	s.Release(-2)
	if got := s.Available(); got != 0 {
		t.Errorf("Available() after no-op releases = %d, want 0", got)
	}
	s.Release(1)
	if got := s.Available(); got != 1 {
		t.Errorf("Available() = %d, want 1", got)
	}
}

func TestStockNoOversell(t *testing.T) {
	const units = 500
	s := NewStock(units)

	var granted atomic.Int64
	const workers, attempts = 16, 100

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < attempts; j++ {
				if s.Reserve(1) {
					granted.Add(1)
				}
			}
		}()
	}
	wg.Wait()

	if got := granted.Load(); got != units {
		t.Errorf("granted = %d, want exactly %d (workers requested %d)", got, units, workers*attempts)
	}
	if got := s.Available(); got != 0 {
		t.Errorf("Available() = %d, want 0", got)
	}
}

func TestStockReserveReleaseBalances(t *testing.T) {
	s := NewStock(8)
	const workers, rounds = 8, 200

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < rounds; j++ {
				if s.Reserve(2) {
					s.Release(2)
				}
			}
		}()
	}
	wg.Wait()

	if got := s.Available(); got != 8 {
		t.Errorf("Available() = %d, want 8", got)
	}
}
