package shardedset

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestAdd(t *testing.T) {
	s := NewSet(4)
	if !s.Add("a") {
		t.Error("the first Add reported false, want true")
	}
	if s.Add("a") {
		t.Error("the second Add reported true, want false")
	}
	if !s.Has("a") {
		t.Error("Has(a) = false, want true")
	}
	if s.Has("b") {
		t.Error("Has(b) = true, want false")
	}
	if s.Len() != 1 {
		t.Errorf("Len = %d, want 1", s.Len())
	}
}

func TestAddManyKeys(t *testing.T) {
	s := NewSet(8)
	for i := 0; i < 1000; i++ {
		if !s.Add(strconv.Itoa(i)) {
			t.Fatalf("Add(%d) reported false on a new key", i)
		}
	}
	if s.Len() != 1000 {
		t.Errorf("Len = %d, want 1000", s.Len())
	}
}

func TestSingleShard(t *testing.T) {
	s := NewSet(0)
	if !s.Add("a") || s.Add("a") {
		t.Error("a single-shard set must still deduplicate")
	}
}

func TestAddIsExactlyOnceUnderConcurrency(t *testing.T) {
	const (
		workers = 16
		keys    = 200
	)
	s := NewSet(16)
	var added atomic.Int64
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < keys; i++ {
				if s.Add(strconv.Itoa(i)) {
					added.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	if got := added.Load(); got != keys {
		t.Errorf("Add reported true %d times, want %d: the check and the insert must be atomic", got, keys)
	}
	if s.Len() != keys {
		t.Errorf("Len = %d, want %d", s.Len(), keys)
	}
}

func TestBucketsDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(bucket{}); got < 64 {
		t.Errorf("sizeof(bucket) = %d, want at least 64", got)
	}
}
