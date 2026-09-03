package batchflush

import (
	"errors"
	"sync"
	"testing"
)

func TestAddFlushesAtTheLimit(t *testing.T) {
	var got [][]int
	b := NewBatcher(2, func(batch []int) error {
		got = append(got, batch)
		return nil
	})
	for i := 1; i <= 4; i++ {
		if err := b.Add(i); err != nil {
			t.Fatal(err)
		}
	}
	if len(got) != 2 {
		t.Fatalf("flushed %d batches, want 2", len(got))
	}
	if got[0][0] != 1 || got[0][1] != 2 || got[1][0] != 3 || got[1][1] != 4 {
		t.Errorf("batches = %v, want [[1 2] [3 4]]", got)
	}
	if b.Pending() != 0 {
		t.Errorf("Pending = %d, want 0", b.Pending())
	}
}

func TestPendingNeverExceedsTheLimit(t *testing.T) {
	b := NewBatcher(4, func([]int) error { return nil })
	for i := 0; i < 1000; i++ {
		if err := b.Add(i); err != nil {
			t.Fatal(err)
		}
		if p := b.Pending(); p >= 4 {
			t.Fatalf("Pending = %d after %d adds, want under 4", p, i+1)
		}
	}
}

func TestBatchesAreIndependentOfThePending(t *testing.T) {
	var kept []int
	b := NewBatcher(2, func(batch []int) error {
		kept = batch
		return nil
	})
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	if kept == nil {
		t.Fatal("no batch was kept")
	}
	if kept[0] != 3 || kept[1] != 4 {
		t.Errorf("the kept batch = %v, want [3 4]: it must not alias the pending buffer", kept)
	}
}

func TestAddPropagatesFlushErrors(t *testing.T) {
	boom := errors.New("boom")
	b := NewBatcher(1, func([]int) error { return boom })
	if err := b.Add(1); !errors.Is(err, boom) {
		t.Errorf("Add = %v, want boom", err)
	}
}

func TestCloseFlushesTheRemainder(t *testing.T) {
	var got [][]int
	b := NewBatcher(3, func(batch []int) error {
		got = append(got, batch)
		return nil
	})
	b.Add(1)
	b.Add(2)
	if err := b.Close(); err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || len(got[0]) != 2 {
		t.Errorf("batches = %v, want one batch of 2", got)
	}
}

func TestConcurrentAdds(t *testing.T) {
	var mu sync.Mutex
	total := 0
	b := NewBatcher(8, func(batch []int) error {
		mu.Lock()
		total += len(batch)
		mu.Unlock()
		return nil
	})
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				b.Add(i)
			}
		}()
	}
	wg.Wait()
	b.Close()
	mu.Lock()
	defer mu.Unlock()
	if total != workers*100 {
		t.Errorf("flushed %d values, want %d", total, workers*100)
	}
}
