package pipeline

import (
	"runtime"
	"testing"
	"time"
)

func feed(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestStageDoublesEverything(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(1, 2, 3), 2)
	sum, count := 0, 0
	for v := range out {
		sum += v
		count++
	}
	if count != 3 || sum != 12 {
		t.Errorf("got %d values summing to %d, want 3 and 12", count, sum)
	}
}

func TestStageClosesOnDrain(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(1), 4)
	<-out
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced an extra value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed after the input drained")
	}
}

func TestStageEmptyInput(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(), 3)
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced a value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed for an empty input")
	}
}

func TestStageZeroWorkers(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(5), 0)
	if v, ok := <-out; !ok || v != 10 {
		t.Errorf("got %d, %v, want 10, true", v, ok)
	}
}

func TestStageDoesNotLeak(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()

	for round := 0; round < 20; round++ {
		done := make(chan struct{})
		in := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case in <- i:
				case <-done:
					close(in)
					return
				}
			}
		}()
		out := Stage(done, in, 4)
		<-out // take one value, then abandon the stage
		close(done)
	}

	deadline := time.Now().Add(3 * time.Second)
	for runtime.NumGoroutine() > base+4 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+4 {
		t.Errorf("goroutines = %d, want about %d: workers are still blocked", got, base)
	}
}
