package fanin

import (
	"runtime"
	"testing"
	"time"
)

func send(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestMergeDeliversEverything(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Merge(done, send(1, 2), send(3), send())
	sum := 0
	count := 0
	for v := range out {
		sum += v
		count++
	}
	if count != 3 || sum != 6 {
		t.Errorf("got %d values summing to %d, want 3 and 6", count, sum)
	}
}

func TestMergeClosesWhenInputsDrain(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Merge(done, send(1))
	<-out
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced an extra value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed after the inputs drained")
	}
}

func TestMergeNoInputs(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Merge(done)
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced a value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed with no inputs")
	}
}

func TestMergeAbandonedConsumerDoesNotLeak(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()

	for round := 0; round < 20; round++ {
		done := make(chan struct{})
		ins := make([]<-chan int, 4)
		for i := range ins {
			ch := make(chan int)
			go func(ch chan int, i int) {
				for j := 0; ; j++ {
					select {
					case ch <- i*100 + j:
					case <-done:
						close(ch)
						return
					}
				}
			}(ch, i)
			ins[i] = ch
		}
		out := Merge(done, ins...)
		<-out // take one value, then walk away
		close(done)
	}

	deadline := time.Now().Add(3 * time.Second)
	for runtime.NumGoroutine() > base+4 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+4 {
		t.Errorf("goroutines = %d, want about %d: forwarders are still blocked on the send", got, base)
	}
}
