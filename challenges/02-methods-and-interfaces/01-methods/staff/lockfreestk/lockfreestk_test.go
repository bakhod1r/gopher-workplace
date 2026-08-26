package lockfreestk

import (
	"sync"
	"testing"
)

func TestLockFreeStack(t *testing.T) {
	s := &Stack{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			s.Push(v)
		}(i)
	}
	wg.Wait()

	count := 0
	curr := s.head.Load()
	for curr != nil {
		count++
		curr = curr.next
	}
	if count != 100 {
		t.Errorf("count = %d, want 100", count)
	}
}
