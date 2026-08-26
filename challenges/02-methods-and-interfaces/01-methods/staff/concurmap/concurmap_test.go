package concurmap

import (
	"sync"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	m := New()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set("k", i)
			m.Get("k")
		}(i)
	}
	wg.Wait()

	if _, ok := m.Get("k"); !ok {
		t.Error("expected key k to exist")
	}
}
