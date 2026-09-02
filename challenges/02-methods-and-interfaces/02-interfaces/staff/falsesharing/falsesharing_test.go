package falsesharing

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestPaddedCellIsOneCacheLine(t *testing.T) {
	if got := unsafe.Sizeof(paddedCell{}); got != CacheLine {
		t.Errorf("sizeof(paddedCell) = %d, want %d", got, CacheLine)
	}
}

func TestAdjacentPaddedCountersAreFarApart(t *testing.T) {
	p := NewPadded(4)
	a := uintptr(unsafe.Pointer(&p.cells[0]))
	b := uintptr(unsafe.Pointer(&p.cells[1]))
	if b-a < CacheLine {
		t.Errorf("adjacent counters are %d bytes apart, want at least %d", b-a, CacheLine)
	}
}

func TestPackedCountersShareLines(t *testing.T) {
	p := NewPacked(4)
	a := uintptr(unsafe.Pointer(&p.vals[0]))
	b := uintptr(unsafe.Pointer(&p.vals[1]))
	if b-a >= CacheLine {
		t.Errorf("packed counters are %d bytes apart; the packed variant should be dense", b-a)
	}
}

func runWorkers(c Counters, workers, iters int) {
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				c.Inc(w)
			}
		}(w)
	}
	wg.Wait()
}

func TestPaddedTotal(t *testing.T) {
	p := NewPadded(4)
	runWorkers(p, 4, 1000)
	if got := p.Total(); got != 4000 {
		t.Errorf("Total = %d, want 4000", got)
	}
}

func TestPackedTotal(t *testing.T) {
	p := NewPacked(4)
	runWorkers(p, 4, 1000)
	if got := p.Total(); got != 4000 {
		t.Errorf("Total = %d, want 4000", got)
	}
}

func TestOutOfRangeIgnored(t *testing.T) {
	p := NewPadded(2)
	p.Inc(-1)
	p.Inc(2)
	p.Inc(100)
	if got := p.Total(); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}

	q := NewPacked(2)
	q.Inc(-1)
	q.Inc(5)
	if got := q.Total(); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
}

func TestBothAreCounters(t *testing.T) {
	var a Counters = NewPadded(1)
	var b Counters = NewPacked(1)
	a.Inc(0)
	b.Inc(0)
	if a.Total() != 1 || b.Total() != 1 {
		t.Errorf("totals = %d, %d; want 1, 1", a.Total(), b.Total())
	}
}

func BenchmarkPaddedParallel(b *testing.B) {
	p := NewPadded(64)
	var slot atomic.Int64
	b.RunParallel(func(pb *testing.PB) {
		i := int(slot.Add(1)) % 64
		for pb.Next() {
			p.Inc(i)
		}
	})
}

func BenchmarkPackedParallel(b *testing.B) {
	p := NewPacked(64)
	var slot atomic.Int64
	b.RunParallel(func(pb *testing.PB) {
		i := int(slot.Add(1)) % 64
		for pb.Next() {
			p.Inc(i)
		}
	})
}
