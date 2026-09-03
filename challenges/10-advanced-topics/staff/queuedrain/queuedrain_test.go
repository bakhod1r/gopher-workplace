package queuedrain

import (
	"runtime"
	"testing"
	"time"
)

func TestQueueTotals(t *testing.T) {
	q := NewQueue(4)
	for i := 1; i <= 100; i++ {
		q.Push(i)
	}
	if got := q.Close(); got != 5050 {
		t.Errorf("Close = %d, want 5050", got)
	}
}

func TestCloseIsIdempotent(t *testing.T) {
	q := NewQueue(2)
	q.Push(7)
	first := q.Close()
	if second := q.Close(); second != first {
		t.Errorf("second Close = %d, want %d", second, first)
	}
}

func TestCloseLeavesNoGoroutines(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()
	for i := 0; i < 20; i++ {
		q := NewQueue(4)
		q.Push(i)
		q.Close()
	}
	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > base+2 && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+2 {
		t.Errorf("goroutines = %d, want about %d: the workers are still blocked on the channel", got, base)
	}
}

func TestCloseWaitsForQueuedWork(t *testing.T) {
	q := NewQueue(1)
	for i := 0; i < 16; i++ {
		q.Push(1)
	}
	if got := q.Close(); got != 16 {
		t.Errorf("Close = %d, want 16: Close returned before the queue drained", got)
	}
}
