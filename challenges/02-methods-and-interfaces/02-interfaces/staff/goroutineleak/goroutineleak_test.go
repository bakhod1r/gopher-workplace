package goroutineleak

import (
	"runtime"
	"testing"
	"time"
)

func settle() {
	for i := 0; i < 50; i++ {
		runtime.Gosched()
		time.Sleep(time.Millisecond)
	}
}

func TestStopExitsGoroutine(t *testing.T) {
	base := runtime.NumGoroutine()

	in := make(chan int)
	w := NewWatcher()
	w.Start(in, &CountingSink{})
	w.Stop()

	settle()
	if got := runtime.NumGoroutine(); got > base {
		t.Errorf("goroutines = %d, baseline %d: the watcher leaked", got, base)
	}
}

func TestClosingInputExitsGoroutine(t *testing.T) {
	base := runtime.NumGoroutine()

	in := make(chan int)
	s := &CountingSink{}
	w := NewWatcher()
	w.Start(in, s)

	in <- 1
	in <- 2
	close(in)
	w.Wait()

	if s.Count() != 2 {
		t.Errorf("Count = %d, want 2", s.Count())
	}

	settle()
	if got := runtime.NumGoroutine(); got > base {
		t.Errorf("goroutines = %d, baseline %d: the watcher leaked", got, base)
	}
}

func TestStopIsIdempotent(t *testing.T) {
	in := make(chan int)
	w := NewWatcher()
	w.Start(in, &CountingSink{})

	w.Stop()
	w.Stop()
	w.Stop()
}

func TestStopWhileBlockedOnSend(t *testing.T) {
	base := runtime.NumGoroutine()

	in := make(chan int) // unbuffered, nobody sends
	w := NewWatcher()
	w.Start(in, &CountingSink{})

	w.Stop() // must not hang

	settle()
	if got := runtime.NumGoroutine(); got > base {
		t.Errorf("goroutines = %d, baseline %d", got, base)
	}
}

func TestManyWatchersDoNotLeak(t *testing.T) {
	base := runtime.NumGoroutine()

	for i := 0; i < 200; i++ {
		in := make(chan int, 1)
		in <- i
		w := NewWatcher()
		w.Start(in, &CountingSink{})
		w.Stop()
	}

	settle()
	if got := runtime.NumGoroutine(); got > base+5 {
		t.Errorf("goroutines = %d, baseline %d: watchers leaked", got, base)
	}
}

func TestObservesBeforeStop(t *testing.T) {
	in := make(chan int)
	s := &CountingSink{}
	w := NewWatcher()
	w.Start(in, s)

	in <- 1
	in <- 2
	w.Stop()

	if s.Count() != 2 {
		t.Errorf("Count = %d, want 2", s.Count())
	}
}
