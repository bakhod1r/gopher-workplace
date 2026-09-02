package workerpool

import (
	"sync/atomic"
	"testing"
)

type countingTask struct {
	n       int
	running *int64
	peak    *int64
}

func (c countingTask) Run() int {
	cur := atomic.AddInt64(c.running, 1)
	for {
		old := atomic.LoadInt64(c.peak)
		if cur <= old || atomic.CompareAndSwapInt64(c.peak, old, cur) {
			break
		}
	}
	atomic.AddInt64(c.running, -1)
	return c.n
}

func TestSquareTask(t *testing.T) {
	if got := (SquareTask{N: 3}).Run(); got != 9 {
		t.Errorf("Run = %d, want 9", got)
	}
}

func TestRunAllOrder(t *testing.T) {
	tasks := make([]Task, 100)
	for i := range tasks {
		tasks[i] = SquareTask{N: i}
	}

	got := RunAll(tasks, 4)
	if len(got) != 100 {
		t.Fatalf("len = %d, want 100", len(got))
	}
	for i := range got {
		if got[i] != i*i {
			t.Fatalf("got[%d] = %d, want %d", i, got[i], i*i)
		}
	}
}

func TestRunAllSingleWorker(t *testing.T) {
	tasks := []Task{SquareTask{2}, SquareTask{3}}
	got := RunAll(tasks, 1)
	if len(got) != 2 || got[0] != 4 || got[1] != 9 {
		t.Errorf("RunAll = %v, want [4 9]", got)
	}
}

func TestRunAllEmpty(t *testing.T) {
	if got := RunAll(nil, 4); len(got) != 0 {
		t.Errorf("RunAll(nil) = %v, want empty", got)
	}
}

func TestConcurrencyBounded(t *testing.T) {
	var running, peak int64
	tasks := make([]Task, 500)
	for i := range tasks {
		tasks[i] = countingTask{n: i, running: &running, peak: &peak}
	}

	RunAll(tasks, 4)
	if peak > 4 {
		t.Errorf("peak concurrency = %d, want at most 4", peak)
	}
}

func BenchmarkRunAll(b *testing.B) {
	tasks := make([]Task, 1000)
	for i := range tasks {
		tasks[i] = SquareTask{N: i}
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RunAll(tasks, 8)
	}
}
