package rwlockopt

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func stores() []Store {
	return []Store{NewRWStore(), NewMutexStore()}
}

func TestGetSetLen(t *testing.T) {
	for _, s := range stores() {
		if _, ok := s.Get("a"); ok {
			t.Errorf("%T: Get on an empty store returned ok", s)
		}
		s.Set("a", 1)
		s.Set("b", 2)
		if v, ok := s.Get("a"); v != 1 || !ok {
			t.Errorf("%T: Get = %d, %v", s, v, ok)
		}
		if s.Len() != 2 {
			t.Errorf("%T: Len = %d, want 2", s, s.Len())
		}
		s.Set("a", 9)
		if v, _ := s.Get("a"); v != 9 {
			t.Errorf("%T: Get after overwrite = %d, want 9", s, v)
		}
	}
}

func TestSnapshotIsIndependent(t *testing.T) {
	s := NewRWStore()
	s.Set("a", 1)

	snap := s.Snapshot()
	s.Set("a", 2)
	s.Set("b", 3)

	if snap["a"] != 1 {
		t.Errorf("snapshot[a] = %d, want 1", snap["a"])
	}
	if _, ok := snap["b"]; ok {
		t.Error("snapshot picked up a later write")
	}
}

func TestConcurrentWriters(t *testing.T) {
	for _, s := range stores() {
		var wg sync.WaitGroup
		const n = 500
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				s.Set(strconv.Itoa(i), i)
			}(i)
		}
		wg.Wait()

		if s.Len() != n {
			t.Errorf("%T: Len = %d, want %d", s, s.Len(), n)
		}
	}
}

func TestReadersOverlap(t *testing.T) {
	if runtime.GOMAXPROCS(0) < 2 {
		t.Skip("needs more than one P")
	}

	s := NewRWStore()
	s.Set("k", 1)

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			deadline := time.Now().Add(200 * time.Millisecond)
			for time.Now().Before(deadline) {
				s.Get("k")
			}
		}()
	}
	wg.Wait()

	if got := s.MaxConcurrentReaders(); got < 2 {
		t.Errorf("max concurrent readers = %d; RLock should let readers overlap", got)
	}
}

func TestConcurrentReadersAndWriters(t *testing.T) {
	for _, s := range stores() {
		var wg sync.WaitGroup
		for i := 0; i < 200; i++ {
			wg.Add(2)
			go func(i int) {
				defer wg.Done()
				s.Set(strconv.Itoa(i%20), i)
			}(i)
			go func(i int) {
				defer wg.Done()
				s.Get(strconv.Itoa(i % 20))
			}(i)
		}
		wg.Wait()
		if s.Len() != 20 {
			t.Errorf("%T: Len = %d, want 20", s, s.Len())
		}
	}
}

func BenchmarkRWGetParallel(b *testing.B) {
	s := NewRWStore()
	s.Set("k", 1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Get("k")
		}
	})
}

func BenchmarkMutexGetParallel(b *testing.B) {
	s := NewMutexStore()
	s.Set("k", 1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Get("k")
		}
	})
}
