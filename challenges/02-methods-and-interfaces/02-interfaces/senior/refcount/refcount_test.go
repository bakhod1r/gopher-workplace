package refcount

import (
	"sync"
	"testing"
)

func TestClosesAtZero(t *testing.T) {
	res := &CountingResource{}
	rc := NewRefCounted(res)

	if !rc.Acquire() {
		t.Fatal("Acquire = false")
	}
	if rc.Count() != 2 {
		t.Errorf("Count = %d, want 2", rc.Count())
	}

	rc.Release()
	if res.ClosedTimes() != 0 {
		t.Error("closed while a reference was still held")
	}
	rc.Release()
	if res.ClosedTimes() != 1 {
		t.Errorf("ClosedTimes = %d, want 1", res.ClosedTimes())
	}
	if rc.Count() != 0 {
		t.Errorf("Count = %d, want 0", rc.Count())
	}
}

func TestReleaseBelowZero(t *testing.T) {
	res := &CountingResource{}
	rc := NewRefCounted(res)

	if !rc.Release() {
		t.Fatal("first Release = false")
	}
	if rc.Release() {
		t.Error("Release below zero should return false")
	}
	if res.ClosedTimes() != 1 {
		t.Errorf("ClosedTimes = %d, want 1", res.ClosedTimes())
	}
}

func TestAcquireAfterClose(t *testing.T) {
	res := &CountingResource{}
	rc := NewRefCounted(res)
	rc.Release()

	if rc.Acquire() {
		t.Error("Acquire after close should return false")
	}
	if rc.Count() != 0 {
		t.Errorf("Count = %d, want 0", rc.Count())
	}
}

func TestConcurrentAcquireRelease(t *testing.T) {
	res := &CountingResource{}
	rc := NewRefCounted(res)

	const n = 500
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rc.Acquire() {
				rc.Release()
			}
		}()
	}
	wg.Wait()

	if res.ClosedTimes() != 0 {
		t.Errorf("closed %d times while the original reference was held", res.ClosedTimes())
	}
	if rc.Count() != 1 {
		t.Errorf("Count = %d, want 1", rc.Count())
	}

	rc.Release()
	if res.ClosedTimes() != 1 {
		t.Errorf("ClosedTimes = %d, want 1", res.ClosedTimes())
	}
}

func TestManyHoldersReleaseConcurrently(t *testing.T) {
	res := &CountingResource{}
	rc := NewRefCounted(res)

	const n = 200
	for i := 0; i < n; i++ {
		if !rc.Acquire() {
			t.Fatal("Acquire = false")
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < n+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rc.Release()
		}()
	}
	wg.Wait()

	if res.ClosedTimes() != 1 {
		t.Errorf("ClosedTimes = %d, want exactly 1", res.ClosedTimes())
	}
}
