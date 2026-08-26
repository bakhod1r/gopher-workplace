package onceinit

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestOnceInit(t *testing.T) {
	var calls int32
	l := New(func() string {
		atomic.AddInt32(&calls, 1)
		return "safe"
	})

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := l.Get(); got != "safe" {
				t.Errorf("got %q, want safe", got)
			}
		}()
	}
	wg.Wait()

	if calls != 1 {
		t.Errorf("calls = %d, want 1", calls)
	}
}
