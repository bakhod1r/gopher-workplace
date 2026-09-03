package deadlinepool

import (
	"context"
	"errors"
	"runtime"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	got, err := Process(context.Background(), []int{1, 2, 3}, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || got[0] != 2 || got[1] != 4 || got[2] != 6 {
		t.Errorf("Process = %v, want [2 4 6]", got)
	}
}

func TestProcessEmpty(t *testing.T) {
	got, err := Process(context.Background(), nil, 4)
	if err != nil || len(got) != 0 {
		t.Errorf("Process = %v, %v, want empty, nil", got, err)
	}
}

func TestProcessWorkerCounts(t *testing.T) {
	items := make([]int, 1001)
	for i := range items {
		items[i] = i
	}
	for _, w := range []int{0, 1, 3, 64, 100000} {
		got, err := Process(context.Background(), items, w)
		if err != nil {
			t.Fatalf("workers=%d: %v", w, err)
		}
		for i := range items {
			if got[i] != items[i]*2 {
				t.Fatalf("workers=%d: got[%d] = %d, want %d", w, i, got[i], items[i]*2)
			}
		}
	}
}

func TestProcessCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := Process(ctx, []int{1, 2, 3}, 2)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}

func TestProcessDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	time.Sleep(20 * time.Millisecond)
	items := make([]int, 10000)
	if _, err := Process(ctx, items, 4); !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want context.DeadlineExceeded", err)
	}
}

func TestProcessDoesNotLeak(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()
	for i := 0; i < 20; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		items := make([]int, 1000)
		Process(ctx, items, 8)
	}
	deadline := time.Now().Add(3 * time.Second)
	for runtime.NumGoroutine() > base+4 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+4 {
		t.Errorf("goroutines = %d, want about %d", got, base)
	}
}
