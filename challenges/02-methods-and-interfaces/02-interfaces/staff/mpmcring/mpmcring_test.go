package mpmcring

import (
	"sync"
	"testing"
)

func TestCapacityRoundsUp(t *testing.T) {
	if got := NewRing(3).Cap(); got != 4 {
		t.Errorf("Cap = %d, want 4", got)
	}
	if got := NewRing(8).Cap(); got != 8 {
		t.Errorf("Cap = %d, want 8", got)
	}
}

func TestFIFO(t *testing.T) {
	r := NewRing(4)
	for i := 1; i <= 3; i++ {
		if !r.Enqueue(i) {
			t.Fatalf("Enqueue(%d) = false", i)
		}
	}
	for i := 1; i <= 3; i++ {
		v, ok := r.Dequeue()
		if !ok || v != i {
			t.Fatalf("Dequeue = %d, %v; want %d, true", v, ok, i)
		}
	}
}

func TestFullAndEmpty(t *testing.T) {
	r := NewRing(2)
	if !r.Enqueue(1) || !r.Enqueue(2) {
		t.Fatal("Enqueue failed below capacity")
	}
	if r.Enqueue(3) {
		t.Error("Enqueue on a full ring should return false")
	}

	r.Dequeue()
	r.Dequeue()
	if _, ok := r.Dequeue(); ok {
		t.Error("Dequeue on an empty ring should return false")
	}
}

func TestWrapAround(t *testing.T) {
	r := NewRing(2)
	for i := 0; i < 100; i++ {
		if !r.Enqueue(i) {
			t.Fatalf("Enqueue(%d) = false", i)
		}
		v, ok := r.Dequeue()
		if !ok || v != i {
			t.Fatalf("Dequeue = %d, %v; want %d", v, ok, i)
		}
	}
}

func TestConcurrentProducersConsumers(t *testing.T) {
	r := NewRing(64)
	const producers, perProducer = 8, 500
	total := producers * perProducer

	var pwg sync.WaitGroup
	for p := 0; p < producers; p++ {
		pwg.Add(1)
		go func(p int) {
			defer pwg.Done()
			for i := 0; i < perProducer; i++ {
				v := p*perProducer + i
				for !r.Enqueue(v) {
					// ring full: retry
				}
			}
		}(p)
	}

	var mu sync.Mutex
	seen := make(map[int]int, total)
	consumed := 0

	var cwg sync.WaitGroup
	for c := 0; c < 8; c++ {
		cwg.Add(1)
		go func() {
			defer cwg.Done()
			for {
				mu.Lock()
				done := consumed >= total
				mu.Unlock()
				if done {
					return
				}

				v, ok := r.Dequeue()
				if !ok {
					continue
				}
				mu.Lock()
				seen[v]++
				consumed++
				mu.Unlock()
			}
		}()
	}

	pwg.Wait()
	cwg.Wait()

	if len(seen) != total {
		t.Fatalf("saw %d distinct values, want %d", len(seen), total)
	}
	for v, c := range seen {
		if c != 1 {
			t.Fatalf("value %d delivered %d times", v, c)
		}
	}
}

func TestIsQueue(t *testing.T) {
	var q Queue = NewRing(2)
	if !q.Enqueue(1) {
		t.Fatal("Enqueue = false")
	}
	if v, ok := q.Dequeue(); v != 1 || !ok {
		t.Errorf("Dequeue = %d, %v", v, ok)
	}
}

func BenchmarkEnqueueDequeue(b *testing.B) {
	r := NewRing(1024)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if r.Enqueue(1) {
				r.Dequeue()
			}
		}
	})
}
