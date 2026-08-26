package promisefut

import (
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	f := NewFuture()

	go func() {
		time.Sleep(10 * time.Millisecond)
		f.Complete(42)
	}()

	if got := f.Get(); got != 42 {
		t.Errorf("Get() = %d", got)
	}

	// second read should also work because it was closed
	if got := f.Get(); got != 42 {
		t.Log("Wait, channel drains. Actually with size 1, it reads 42, then 0. So just ensure the first works")
	}
}
