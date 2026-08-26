package shardedmap

import (
	"sync"
	"testing"
)

func TestShardedMap(t *testing.T) {
	m := New(32)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set("key", i)
		}(i)
	}
	wg.Wait()

	// verify it didn't panic
}
