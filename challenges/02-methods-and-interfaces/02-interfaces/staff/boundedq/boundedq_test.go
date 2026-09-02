package boundedq

import (
	"sync"
	"testing"
)

func TestPushPop(t *testing.T) {
	q := NewQueue(2)
	if !q.Push(1) || !q.Push(2) {
		t.Fatal("Push returned false on an open queue")
	}
	if q.Len() != 2 {
		t.Errorf("Len = %d, want 2", q.Len())
	}
	if v, ok := q.Pop(); v != 1 || !ok {
		t.Errorf("Pop = %d, %v; want 1, true (FIFO)", v, ok)
	}
	if v, ok := q.Pop(); v != 2 || !ok {
		t.Errorf("Pop = %d, %v; want 2, true", v, ok)
	}
}

func TestPushAfterClose(t *testing.T) {
	q := NewQueue(2)
	q.Close()
	if q.Push(1) {
		t.Error("Push after Close should return false")
	}
}

func TestPopDrainsThenReportsClosed(t *testing.T) {
	q := NewQueue(4)
	q.Push(1)
	q.Push(2)
	q.Close()

	if v, ok := q.Pop(); v != 1 || !ok {
		t.Errorf("Pop = %d, %v; want 1, true", v, ok)
	}
	if v, ok := q.Pop(); v != 2 || !ok {
		t.Errorf("Pop = %d, %v; want 2, true", v, ok)
	}
	if _, ok := q.Pop(); ok {
		t.Error("Pop on a drained closed queue should return false")
	}
}

func TestCloseWakesBlockedConsumers(t *testing.T) {
	q := NewQueue(1)
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, ok := q.Pop(); ok {
				t.Error("Pop returned an item from an empty closed queue")
			}
		}()
	}

	q.Close()
	wg.Wait()
}

func TestCloseWakesBlockedProducers(t *testing.T) {
	q := NewQueue(1)
	q.Push(1) // queue is now full

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if q.Push(2) {
				t.Error("Push should fail once the queue is closed")
			}
		}()
	}

	q.Close()
	wg.Wait()
}

func TestBackpressure(t *testing.T) {
	q := NewQueue(4)
	const n = 2000

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			if !q.Push(i) {
				t.Error("Push failed on an open queue")
				return
			}
			if q.Len() > 4 {
				t.Errorf("Len = %d, exceeded the capacity", q.Len())
				return
			}
		}
		q.Close()
	}()

	got := 0
	for {
		_, ok := q.Pop()
		if !ok {
			break
		}
		got++
	}
	wg.Wait()

	if got != n {
		t.Errorf("consumed %d items, want %d", got, n)
	}
}

func TestMultipleProducersConsumers(t *testing.T) {
	q := NewQueue(8)
	const producers, perProducer = 4, 250

	var pwg sync.WaitGroup
	for p := 0; p < producers; p++ {
		pwg.Add(1)
		go func() {
			defer pwg.Done()
			for i := 0; i < perProducer; i++ {
				q.Push(i)
			}
		}()
	}

	var mu sync.Mutex
	total := 0
	var cwg sync.WaitGroup
	for c := 0; c < 4; c++ {
		cwg.Add(1)
		go func() {
			defer cwg.Done()
			n := 0
			for {
				if _, ok := q.Pop(); !ok {
					break
				}
				n++
			}
			mu.Lock()
			total += n
			mu.Unlock()
		}()
	}

	pwg.Wait()
	q.Close()
	cwg.Wait()

	if total != producers*perProducer {
		t.Errorf("consumed %d items, want %d", total, producers*perProducer)
	}
}

func TestIsSink(t *testing.T) {
	var s Sink = NewQueue(1)
	if !s.Push(1) {
		t.Error("Push = false")
	}
}
