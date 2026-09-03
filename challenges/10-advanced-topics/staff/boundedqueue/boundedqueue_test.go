package boundedqueue

import (
	"sync"
	"testing"
	"time"
)

func TestPutAndTake(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	q := NewQueue(2)
	if !q.Put(done, 1) || !q.Put(done, 2) {
		t.Fatal("the first two puts must succeed")
	}
	if v, ok := q.Take(done); !ok || v != 1 {
		t.Errorf("Take = %d, %v, want 1, true", v, ok)
	}
	if v, ok := q.Take(done); !ok || v != 2 {
		t.Errorf("Take = %d, %v, want 2, true: the queue must be FIFO", v, ok)
	}
}

func TestPutBlocksWhenFull(t *testing.T) {
	done := make(chan struct{})
	q := NewQueue(1)
	q.Put(done, 1)
	got := make(chan bool, 1)
	go func() { got <- q.Put(done, 2) }()
	select {
	case <-got:
		t.Fatal("Put returned while the queue was full")
	case <-time.After(50 * time.Millisecond):
	}
	close(done)
	if ok := <-got; ok {
		t.Error("the cancelled Put reported true, want false")
	}
}

func TestPutResumesAfterTake(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	q := NewQueue(1)
	q.Put(done, 1)
	got := make(chan bool, 1)
	go func() { got <- q.Put(done, 2) }()
	time.Sleep(20 * time.Millisecond)
	if v, _ := q.Take(done); v != 1 {
		t.Fatalf("Take = %d, want 1", v)
	}
	select {
	case ok := <-got:
		if !ok {
			t.Error("the waiting Put reported false, want true")
		}
	case <-time.After(2 * time.Second):
		t.Error("the waiting Put was never released")
	}
}

func TestQueueNeverExceedsCapacity(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	q := NewQueue(4)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if !q.Put(done, i) {
				return
			}
			if q.Len() > q.Cap() {
				panic("the queue grew past its capacity")
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if _, ok := q.Take(done); !ok {
				return
			}
		}
	}()
	wg.Wait()
}
