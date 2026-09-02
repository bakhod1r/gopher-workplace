package fanin

import (
	"runtime"
	"testing"
	"time"
)

func filled(vs ...int) <-chan int {
	ch := make(chan int, len(vs))
	for _, v := range vs {
		ch <- v
	}
	close(ch)
	return ch
}

func settle() {
	for i := 0; i < 50; i++ {
		runtime.Gosched()
		time.Sleep(time.Millisecond)
	}
}

func TestMergeDrains(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	out := Fan{}.Merge(done, filled(1, 2), filled(3))

	sum, n := 0, 0
	for v := range out {
		sum += v
		n++
	}
	if n != 3 || sum != 6 {
		t.Errorf("n = %d, sum = %d; want 3, 6", n, sum)
	}
}

func TestMergeNoInputs(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	out := Fan{}.Merge(done)
	select {
	case _, ok := <-out:
		if ok {
			t.Error("expected a closed output")
		}
	case <-time.After(time.Second):
		t.Fatal("the output channel was never closed")
	}
}

func TestOutputClosesAfterInputs(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	out := Fan{}.Merge(done, filled(1), filled(2))
	for range out {
	}

	select {
	case _, ok := <-out:
		if ok {
			t.Error("expected a closed output")
		}
	default:
		t.Error("the output should already be closed")
	}
}

func TestEarlyExitDoesNotLeak(t *testing.T) {
	base := runtime.NumGoroutine()

	// Producers that never stop on their own.
	mk := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case ch <- i:
				case <-time.After(3 * time.Second):
					return
				}
			}
		}()
		return ch
	}

	done := make(chan struct{})
	out := Fan{}.Merge(done, mk(), mk(), mk())

	<-out // read exactly one value, then walk away
	close(done)

	settle()
	if got := runtime.NumGoroutine(); got > base+4 {
		t.Errorf("goroutines = %d, baseline %d: forwarders leaked", got, base)
	}
}

func TestDoneClosesOutput(t *testing.T) {
	done := make(chan struct{})
	src := make(chan int)
	out := Fan{}.Merge(done, src)

	close(done)

	select {
	case _, ok := <-out:
		if ok {
			t.Error("expected the output to be closed after done")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("the output was never closed after done")
	}
}

func TestManyInputs(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	ins := make([]<-chan int, 0, 20)
	for i := 0; i < 20; i++ {
		ins = append(ins, filled(i))
	}

	out := Fan{}.Merge(done, ins...)
	n := 0
	for range out {
		n++
	}
	if n != 20 {
		t.Errorf("received %d values, want 20", n)
	}
}

func TestIsMerger(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	var m Merger = Fan{}
	out := m.Merge(done, filled(7))
	if v, ok := <-out; v != 7 || !ok {
		t.Errorf("got %d, %v; want 7, true", v, ok)
	}
}
