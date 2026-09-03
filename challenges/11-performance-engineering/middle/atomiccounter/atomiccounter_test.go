package atomiccounter

import (
	"sync"
	"testing"
)

func TestAddReturnsTheNewValue(t *testing.T) {
	var s Stats
	if got := s.Add(3); got != 3 {
		t.Errorf("Add = %d, want 3", got)
	}
	if got := s.Add(-1); got != 2 {
		t.Errorf("Add = %d, want 2", got)
	}
	if got := s.Total(); got != 2 {
		t.Errorf("Total = %d, want 2", got)
	}
}

func TestObserveKeepsTheMaximum(t *testing.T) {
	var s Stats
	if got := s.Max(); got != 0 {
		t.Errorf("Max = %d, want 0", got)
	}
	s.Observe(5)
	s.Observe(2)
	s.Observe(4)
	if got := s.Max(); got != 5 {
		t.Errorf("Max = %d, want 5", got)
	}
	s.Observe(9)
	if got := s.Max(); got != 9 {
		t.Errorf("Max = %d, want 9", got)
	}
}

func TestConcurrentAdds(t *testing.T) {
	var s Stats
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				s.Add(1)
			}
		}()
	}
	wg.Wait()
	if got := s.Total(); got != 50000 {
		t.Errorf("Total = %d, want 50000", got)
	}
}

func TestConcurrentObserve(t *testing.T) {
	var s Stats
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 100; j++ {
				s.Observe(int64(i * j))
			}
		}()
	}
	wg.Wait()
	if got := s.Max(); got != 10000 {
		t.Errorf("Max = %d, want 10000 — a lost compare-and-swap drops the true maximum", got)
	}
}

func TestObserveIgnoresSmallerValuesUnderContention(t *testing.T) {
	var s Stats
	s.Observe(1000)
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				s.Observe(int64(j))
			}
		}()
	}
	wg.Wait()
	if got := s.Max(); got != 1000 {
		t.Errorf("Max = %d, want 1000 — smaller observations must never overwrite the maximum", got)
	}
}
