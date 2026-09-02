package semaphoreifc

import (
	"sync/atomic"
	"testing"
)

type peakJob struct {
	running *int64
	peak    *int64
	done    *int64
}

func (p peakJob) Do() {
	cur := atomic.AddInt64(p.running, 1)
	for {
		old := atomic.LoadInt64(p.peak)
		if cur <= old || atomic.CompareAndSwapInt64(p.peak, old, cur) {
			break
		}
	}
	atomic.AddInt64(p.done, 1)
	atomic.AddInt64(p.running, -1)
}

func TestTryAcquire(t *testing.T) {
	s := NewSemaphore(1)
	if !s.TryAcquire() {
		t.Fatal("first TryAcquire should succeed")
	}
	if s.TryAcquire() {
		t.Error("second TryAcquire should fail")
	}
	s.Release()
	if !s.TryAcquire() {
		t.Error("TryAcquire after Release should succeed")
	}
}

func TestAcquireRelease(t *testing.T) {
	s := NewSemaphore(2)
	s.Acquire()
	s.Acquire()
	if s.TryAcquire() {
		t.Error("semaphore should be full")
	}
	s.Release()
	if !s.TryAcquire() {
		t.Error("a slot should be free")
	}
}

func TestIsLimiter(t *testing.T) {
	var l Limiter = NewSemaphore(1)
	l.Acquire()
	l.Release()
}

func TestRunLimitedBounds(t *testing.T) {
	var running, peak, done int64
	jobs := make([]Job, 500)
	for i := range jobs {
		jobs[i] = peakJob{running: &running, peak: &peak, done: &done}
	}

	RunLimited(jobs, 3)

	if peak > 3 {
		t.Errorf("peak concurrency = %d, want at most 3", peak)
	}
	if done != 500 {
		t.Errorf("completed %d jobs, want 500", done)
	}
	if running != 0 {
		t.Errorf("%d jobs still running after RunLimited returned", running)
	}
}

func TestRunLimitedEmpty(t *testing.T) {
	RunLimited(nil, 4)
}

func TestRunLimitedSingleSlot(t *testing.T) {
	var running, peak, done int64
	jobs := make([]Job, 50)
	for i := range jobs {
		jobs[i] = peakJob{running: &running, peak: &peak, done: &done}
	}

	RunLimited(jobs, 1)
	if peak > 1 {
		t.Errorf("peak = %d, want 1", peak)
	}
	if done != 50 {
		t.Errorf("done = %d, want 50", done)
	}
}
