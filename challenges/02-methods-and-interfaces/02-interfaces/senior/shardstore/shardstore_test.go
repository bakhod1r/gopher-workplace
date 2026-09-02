package shardstore

import (
	"strconv"
	"sync"
	"testing"
)

func TestPutGet(t *testing.T) {
	s := NewShardedStore(4)
	s.Put("a", "1")
	if v, ok := s.Get("a"); v != "1" || !ok {
		t.Errorf("Get = %q, %v", v, ok)
	}
	if _, ok := s.Get("zz"); ok {
		t.Error("Get on a missing key returned ok")
	}
	s.Put("a", "2")
	if v, _ := s.Get("a"); v != "2" {
		t.Errorf("Get after overwrite = %q, want \"2\"", v)
	}
}

func TestStableSharding(t *testing.T) {
	s := NewShardedStore(8)
	for i := 0; i < 1000; i++ {
		k := strconv.Itoa(i)
		if s.shardFor(k) != s.shardFor(k) {
			t.Fatalf("key %q mapped to two different shards", k)
		}
	}
}

func TestSingleShard(t *testing.T) {
	s := NewShardedStore(0) // clamped to 1
	s.Put("a", "1")
	if v, ok := s.Get("a"); v != "1" || !ok {
		t.Errorf("Get = %q, %v", v, ok)
	}
	if s.Len() != 1 {
		t.Errorf("Len = %d, want 1", s.Len())
	}
}

func TestConcurrentWriters(t *testing.T) {
	s := NewShardedStore(16)
	const n = 1000

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			k := strconv.Itoa(i)
			s.Put(k, k)
		}(i)
	}
	wg.Wait()

	if s.Len() != n {
		t.Fatalf("Len = %d, want %d", s.Len(), n)
	}
	for i := 0; i < n; i++ {
		k := strconv.Itoa(i)
		if v, ok := s.Get(k); !ok || v != k {
			t.Fatalf("Get(%q) = %q, %v", k, v, ok)
		}
	}
}

func TestConcurrentReadersAndWriters(t *testing.T) {
	s := NewShardedStore(8)
	var wg sync.WaitGroup

	for i := 0; i < 200; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			s.Put(strconv.Itoa(i), "v")
		}(i)
		go func(i int) {
			defer wg.Done()
			s.Get(strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
}

func TestIsStore(t *testing.T) {
	var st Store = NewShardedStore(4)
	st.Put("k", "v")
	if v, ok := st.Get("k"); v != "v" || !ok {
		t.Errorf("Get = %q, %v", v, ok)
	}
}

func BenchmarkPutParallel(b *testing.B) {
	s := NewShardedStore(32)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			s.Put(strconv.Itoa(i%1000), "v")
			i++
		}
	})
}
