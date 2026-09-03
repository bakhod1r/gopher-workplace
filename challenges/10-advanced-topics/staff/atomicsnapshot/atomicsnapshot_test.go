package atomicsnapshot

import (
	"sync"
	"testing"
)

func TestSetThenGet(t *testing.T) {
	var s Store
	if got := s.Get(); got != (Config{}) {
		t.Errorf("Get = %v, want the zero Config", got)
	}
	s.Set(Config{Retries: 3, Timeout: 10})
	if got := s.Get(); got != (Config{Retries: 3, Timeout: 10}) {
		t.Errorf("Get = %v, want {3 10}", got)
	}
	s.Set(Config{Retries: 1})
	if got := s.Get(); got.Retries != 1 || got.Timeout != 0 {
		t.Errorf("Get = %v, want {1 0}", got)
	}
}

func TestSetDoesNotAliasTheCaller(t *testing.T) {
	var s Store
	c := Config{Retries: 1, Timeout: 2}
	s.Set(c)
	c.Retries = 99
	if got := s.Get(); got.Retries != 1 {
		t.Errorf("Retries = %d, want 1: the store aliases the caller's variable", got.Retries)
	}
}

func TestSnapshotsNeverTear(t *testing.T) {
	var s Store
	s.Set(Config{Retries: 1, Timeout: 10})
	var wg sync.WaitGroup
	stop := make(chan struct{})
	for w := 0; w < 4; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
				}
				c := s.Get()
				if c.Timeout != c.Retries*10 {
					panic("torn snapshot")
				}
			}
		}()
	}
	for i := 1; i <= 2000; i++ {
		s.Set(Config{Retries: i, Timeout: i * 10})
	}
	close(stop)
	wg.Wait()
}
