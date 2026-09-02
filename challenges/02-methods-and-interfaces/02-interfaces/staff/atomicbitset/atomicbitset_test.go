package atomicbitset

import (
	"sync"
	"testing"
)

func TestSetClearTest(t *testing.T) {
	b := NewBitset(128)

	if !b.Set(5) {
		t.Error("first Set should report a change")
	}
	if b.Set(5) {
		t.Error("second Set should report no change")
	}
	if !b.Test(5) {
		t.Error("Test = false after Set")
	}
	if b.Count() != 1 {
		t.Errorf("Count = %d, want 1", b.Count())
	}

	if !b.Clear(5) {
		t.Error("first Clear should report a change")
	}
	if b.Clear(5) {
		t.Error("second Clear should report no change")
	}
	if b.Test(5) {
		t.Error("Test = true after Clear")
	}
	if b.Count() != 0 {
		t.Errorf("Count = %d, want 0", b.Count())
	}
}

func TestWordBoundaries(t *testing.T) {
	b := NewBitset(200)
	for _, i := range []int{0, 63, 64, 127, 128, 199} {
		if !b.Set(i) {
			t.Fatalf("Set(%d) = false", i)
		}
		if !b.Test(i) {
			t.Fatalf("Test(%d) = false", i)
		}
	}
	if b.Count() != 6 {
		t.Errorf("Count = %d, want 6", b.Count())
	}
}

func TestOutOfRange(t *testing.T) {
	b := NewBitset(64)
	if b.Set(-1) || b.Set(64) || b.Set(1000) {
		t.Error("out-of-range Set should report no change")
	}
	if b.Test(-1) || b.Test(64) {
		t.Error("out-of-range Test should be false")
	}
	if b.Clear(999) {
		t.Error("out-of-range Clear should report no change")
	}
	if b.Count() != 0 {
		t.Errorf("Count = %d, want 0", b.Count())
	}
}

func TestConcurrentSetsInOneWord(t *testing.T) {
	b := NewBitset(64) // every bit shares one word

	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if !b.Set(i) {
				t.Errorf("Set(%d) reported no change", i)
			}
		}(i)
	}
	wg.Wait()

	if got := b.Count(); got != 64 {
		t.Errorf("Count = %d, want 64 (lost updates)", got)
	}
	for i := 0; i < 64; i++ {
		if !b.Test(i) {
			t.Fatalf("bit %d was lost", i)
		}
	}
}

func TestConcurrentSetSameBitReportsOneChange(t *testing.T) {
	b := NewBitset(64)

	var mu sync.Mutex
	changes := 0
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if b.Set(7) {
				mu.Lock()
				changes++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if changes != 1 {
		t.Errorf("%d goroutines reported a change, want exactly 1", changes)
	}
}

func TestConcurrentSetAndClear(t *testing.T) {
	b := NewBitset(1024)
	var wg sync.WaitGroup

	for i := 0; i < 1024; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			b.Set(i)
		}(i)
		go func(i int) {
			defer wg.Done()
			b.Clear(i)
		}(i)
	}
	wg.Wait()

	// The final state is racy by design; it just has to be consistent.
	n := 0
	for i := 0; i < 1024; i++ {
		if b.Test(i) {
			n++
		}
	}
	if n != b.Count() {
		t.Errorf("Count = %d, but %d bits test set", b.Count(), n)
	}
}

func TestIsSet(t *testing.T) {
	var s Set = NewBitset(8)
	s.Set(1)
	if !s.Test(1) || s.Count() != 1 {
		t.Errorf("Test = %v, Count = %d", s.Test(1), s.Count())
	}
}

func BenchmarkSetParallel(b *testing.B) {
	bs := NewBitset(1 << 16)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			bs.Set(i % (1 << 16))
			i++
		}
	})
}
