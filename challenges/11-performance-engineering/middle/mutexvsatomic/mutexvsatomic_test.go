package mutexvsatomic

import (
	"sync"
	"testing"
)

func TestMutexCounter(t *testing.T) {
	var c MutexCounter
	c.Inc()
	c.Inc()
	if got := c.Value(); got != 2 {
		t.Errorf("Value = %d, want 2", got)
	}
}

func TestAtomicCounter(t *testing.T) {
	var c AtomicCounter
	c.Inc()
	c.Inc()
	if got := c.Value(); got != 2 {
		t.Errorf("Value = %d, want 2", got)
	}
}

func TestShardedCounter(t *testing.T) {
	c := NewSharded(4)
	c.Inc(0)
	c.Inc(1)
	c.Inc(9)
	if got := c.Value(); got != 3 {
		t.Errorf("Value = %d, want 3", got)
	}
}

func TestShardedNonPositive(t *testing.T) {
	for _, n := range []int{0, -3} {
		c := NewSharded(n)
		c.Inc(5)
		if got := c.Value(); got != 1 {
			t.Errorf("NewSharded(%d): Value = %d, want 1", n, got)
		}
	}
}

func TestShardedNegativeIDDoesNotPanic(t *testing.T) {
	c := NewSharded(4)
	c.Inc(-7)
	if got := c.Value(); got != 1 {
		t.Errorf("Value = %d, want 1", got)
	}
}

func TestAllThreeAgreeUnderConcurrency(t *testing.T) {
	const goroutines, each = 50, 2000
	var m MutexCounter
	var a AtomicCounter
	s := NewSharded(16)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < each; j++ {
				m.Inc()
				a.Inc()
				s.Inc(i)
			}
		}()
	}
	wg.Wait()

	want := int64(goroutines * each)
	if got := m.Value(); got != want {
		t.Errorf("MutexCounter = %d, want %d", got, want)
	}
	if got := a.Value(); got != want {
		t.Errorf("AtomicCounter = %d, want %d", got, want)
	}
	if got := s.Value(); got != want {
		t.Errorf("ShardedCounter = %d, want %d", got, want)
	}
}
