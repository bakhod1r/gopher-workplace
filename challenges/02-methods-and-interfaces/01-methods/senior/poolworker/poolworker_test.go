package poolworker

import (
	"sync/atomic"
	"testing"
)

func TestPool(t *testing.T) {
	p := Pool{
		Count: 3,
		Tasks: make(chan func(), 10),
	}
	p.Start()

	var sum int32
	for i := 0; i < 5; i++ {
		p.Tasks <- func() { atomic.AddInt32(&sum, 1) }
	}
	close(p.Tasks)
	p.Wait()

	if sum != 5 {
		t.Errorf("sum = %d, want 5", sum)
	}
}
