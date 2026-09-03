package copyonwrite

import (
	"sync"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	var s Store
	if _, ok := s.Get("a"); ok {
		t.Error("Get on an empty store reported ok, want false")
	}
	s.Set("a", 1)
	s.Set("b", 2)
	if v, ok := s.Get("a"); !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if v, ok := s.Get("b"); !ok || v != 2 {
		t.Errorf("Get(b) = %d, %v, want 2, true", v, ok)
	}
	if s.Len() != 2 {
		t.Errorf("Len = %d, want 2", s.Len())
	}
}

func TestSetOverwrites(t *testing.T) {
	var s Store
	s.Set("a", 1)
	s.Set("a", 9)
	if v, _ := s.Get("a"); v != 9 {
		t.Errorf("Get(a) = %d, want 9", v)
	}
	if s.Len() != 1 {
		t.Errorf("Len = %d, want 1", s.Len())
	}
}

func TestOldSnapshotsAreImmutable(t *testing.T) {
	var s Store
	s.Set("a", 1)
	before := s.m.Load()
	s.Set("b", 2)
	if len(*before) != 1 {
		t.Errorf("the previous snapshot grew to %d entries: it was modified in place", len(*before))
	}
	if _, ok := (*before)["b"]; ok {
		t.Error("the previous snapshot gained a key")
	}
}

func TestConcurrentReadersAndWriters(t *testing.T) {
	var s Store
	s.Set("seed", 0)
	var wg sync.WaitGroup
	stop := make(chan struct{})

	for r := 0; r < 8; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
				}
				if v, ok := s.Get("seed"); !ok || v != 0 {
					panic("seed changed")
				}
			}
		}()
	}

	var wwg sync.WaitGroup
	for w := 0; w < 4; w++ {
		wwg.Add(1)
		go func(w int) {
			defer wwg.Done()
			for i := 0; i < 200; i++ {
				s.Set(string(rune('a'+w)), i)
			}
		}(w)
	}
	wwg.Wait()
	close(stop)
	wg.Wait()

	if s.Len() != 5 {
		t.Errorf("Len = %d, want 5", s.Len())
	}
}
