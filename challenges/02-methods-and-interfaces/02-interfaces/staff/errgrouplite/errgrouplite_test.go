package errgrouplite

import (
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

var errBoom = errors.New("boom")

func ok() Task {
	return TaskFunc(func(<-chan struct{}) error { return nil })
}

func fail(err error) Task {
	return TaskFunc(func(<-chan struct{}) error { return err })
}

func TestAllSucceed(t *testing.T) {
	g := NewGroup(4)
	for i := 0; i < 10; i++ {
		g.Go(ok())
	}
	if err := g.Wait(); err != nil {
		t.Errorf("Wait = %v, want nil", err)
	}
}

func TestFirstErrorWins(t *testing.T) {
	g := NewGroup(1) // serialise so the order is deterministic
	g.Go(ok())
	g.Go(fail(errBoom))
	g.Go(fail(errors.New("later")))

	err := g.Wait()
	if !errors.Is(err, errBoom) {
		t.Errorf("Wait = %v, want the first error", err)
	}
}

func TestCancelSignalsRunningTasks(t *testing.T) {
	g := NewGroup(4)

	var cancelled atomic.Int64
	for i := 0; i < 3; i++ {
		g.Go(TaskFunc(func(cancel <-chan struct{}) error {
			select {
			case <-cancel:
				cancelled.Add(1)
				return nil
			case <-time.After(5 * time.Second):
				return errors.New("never cancelled")
			}
		}))
	}
	g.Go(fail(errBoom))

	if err := g.Wait(); !errors.Is(err, errBoom) {
		t.Fatalf("Wait = %v, want errBoom", err)
	}
	if cancelled.Load() != 3 {
		t.Errorf("%d tasks saw the cancellation, want 3", cancelled.Load())
	}
}

func TestConcurrencyBounded(t *testing.T) {
	var running, peak int64
	g := NewGroup(2)

	for i := 0; i < 100; i++ {
		g.Go(TaskFunc(func(<-chan struct{}) error {
			cur := atomic.AddInt64(&running, 1)
			for {
				old := atomic.LoadInt64(&peak)
				if cur <= old || atomic.CompareAndSwapInt64(&peak, old, cur) {
					break
				}
			}
			atomic.AddInt64(&running, -1)
			return nil
		}))
	}

	if err := g.Wait(); err != nil {
		t.Fatalf("Wait = %v", err)
	}
	if peak > 2 {
		t.Errorf("peak concurrency = %d, want at most 2", peak)
	}
}

func TestWaitJoinsEveryGoroutine(t *testing.T) {
	base := runtime.NumGoroutine()

	g := NewGroup(4)
	for i := 0; i < 50; i++ {
		g.Go(ok())
	}
	g.Wait()

	for i := 0; i < 50; i++ {
		runtime.Gosched()
		time.Sleep(time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+2 {
		t.Errorf("goroutines = %d, baseline %d: Wait did not join them", got, base)
	}
}

func TestEmptyGroup(t *testing.T) {
	if err := NewGroup(4).Wait(); err != nil {
		t.Errorf("Wait = %v, want nil", err)
	}
}

func TestCancelledChannelOpenOnSuccess(t *testing.T) {
	g := NewGroup(2)
	g.Go(ok())
	g.Wait()

	select {
	case <-g.Cancelled():
		t.Error("the cancel channel should stay open when nothing failed")
	default:
	}
}
