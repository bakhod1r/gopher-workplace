package atomicvsmutex

import (
	"sync"
	"testing"
)

func counters() []Counter {
	return []Counter{&MutexCounter{}, &AtomicCounter{}}
}

func TestAddAndValue(t *testing.T) {
	for _, c := range counters() {
		c.Add(5)
		c.Add(5)
		c.Add(5)
		if got := c.Value(); got != 15 {
			t.Errorf("%T: Value = %d, want 15", c, got)
		}
	}
}

func TestNegativeDelta(t *testing.T) {
	for _, c := range counters() {
		c.Add(10)
		c.Add(-3)
		if got := c.Value(); got != 7 {
			t.Errorf("%T: Value = %d, want 7", c, got)
		}
	}
}

func TestZeroValueUsable(t *testing.T) {
	var m MutexCounter
	var a AtomicCounter
	m.Inc()
	a.Inc()
	if m.Value() != 1 || a.Value() != 1 {
		t.Errorf("values = %d, %d; want 1, 1", m.Value(), a.Value())
	}
}

func TestNoLostUpdates(t *testing.T) {
	const n = 2000

	for _, c := range counters() {
		var wg sync.WaitGroup
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Inc()
			}()
		}
		wg.Wait()

		if got := c.Value(); got != n {
			t.Errorf("%T: Value = %d, want %d (lost updates)", c, got, n)
		}
	}
}

func TestConcurrentMixedOperations(t *testing.T) {
	for _, c := range counters() {
		var wg sync.WaitGroup
		for i := 0; i < 500; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				c.Add(2)
			}()
			go func() {
				defer wg.Done()
				c.Add(-1)
			}()
		}
		wg.Wait()

		if got := c.Value(); got != 500 {
			t.Errorf("%T: Value = %d, want 500", c, got)
		}
	}
}

func TestIncAll(t *testing.T) {
	a, b := &MutexCounter{}, &AtomicCounter{}
	IncAll(a, b)
	IncAll(a, b)
	if a.Value() != 2 || b.Value() != 2 {
		t.Errorf("values = %d, %d; want 2, 2", a.Value(), b.Value())
	}
	IncAll()
}

func TestAtomicIncDoesNotAllocate(t *testing.T) {
	c := &AtomicCounter{}
	if avg := testing.AllocsPerRun(1000, func() { c.Inc() }); avg > 0 {
		t.Errorf("AtomicCounter.Inc allocated %.2f times per call, want 0", avg)
	}
}

func BenchmarkMutexIncParallel(b *testing.B) {
	c := &MutexCounter{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}

func BenchmarkAtomicIncParallel(b *testing.B) {
	c := &AtomicCounter{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}
