package batchedcounter

import (
	"sync"
	"testing"
)

func TestAddAccumulatesLocally(t *testing.T) {
	var c Counter
	var l Local
	for i := 0; i < batchSize-1; i++ {
		c.Add(&l, 1)
	}
	if got := c.Total(); got != 0 {
		t.Errorf("Total = %d, want 0: nothing should have been published yet", got)
	}
	c.Flush(&l)
	if got := c.Total(); got != batchSize-1 {
		t.Errorf("Total = %d, want %d", got, batchSize-1)
	}
}

func TestAddPublishesAtTheThreshold(t *testing.T) {
	var c Counter
	var l Local
	for i := 0; i < batchSize; i++ {
		c.Add(&l, 1)
	}
	if got := c.Total(); got != batchSize {
		t.Errorf("Total = %d, want %d: the batch should have been published", got, batchSize)
	}
	if l.n != 0 {
		t.Errorf("local = %d, want 0 after publishing", l.n)
	}
}

func TestAddHandlesLargeIncrements(t *testing.T) {
	var c Counter
	var l Local
	c.Add(&l, 1000)
	if got := c.Total(); got != 1000 {
		t.Errorf("Total = %d, want 1000", got)
	}
}

func TestAddHandlesNegatives(t *testing.T) {
	var c Counter
	var l Local
	c.Add(&l, -1000)
	c.Flush(&l)
	if got := c.Total(); got != -1000 {
		t.Errorf("Total = %d, want -1000", got)
	}
}

func TestNoIncrementIsLost(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	const (
		workers = 16
		each    = 1000
	)
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			var l Local
			for i := 0; i < each; i++ {
				c.Add(&l, 1)
			}
			c.Flush(&l)
		}()
	}
	wg.Wait()
	if got := c.Total(); got != workers*each {
		t.Errorf("Total = %d, want %d", got, workers*each)
	}
}

func BenchmarkAdd(b *testing.B) {
	var c Counter
	b.RunParallel(func(pb *testing.PB) {
		var l Local
		for pb.Next() {
			c.Add(&l, 1)
		}
		c.Flush(&l)
	})
}
