package lockfreestk

import (
	"sync"
	"testing"
)

func TestLIFO(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)

	if v, ok := s.Pop(); v != 2 || !ok {
		t.Errorf("Pop = %d, %v; want 2, true", v, ok)
	}
	if v, ok := s.Pop(); v != 1 || !ok {
		t.Errorf("Pop = %d, %v; want 1, true", v, ok)
	}
	if _, ok := s.Pop(); ok {
		t.Error("Pop on an empty stack returned ok")
	}
}

func TestLen(t *testing.T) {
	var s Stack
	if s.Len() != 0 {
		t.Errorf("Len = %d, want 0", s.Len())
	}
	s.Push(1)
	s.Push(2)
	if s.Len() != 2 {
		t.Errorf("Len = %d, want 2", s.Len())
	}
	s.Pop()
	if s.Len() != 1 {
		t.Errorf("Len = %d, want 1", s.Len())
	}
}

func TestConcurrentPush(t *testing.T) {
	var s Stack
	const n = 1000

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Push(i)
		}(i)
	}
	wg.Wait()

	if s.Len() != n {
		t.Fatalf("Len = %d, want %d", s.Len(), n)
	}

	seen := make(map[int]int, n)
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		seen[v]++
	}
	if len(seen) != n {
		t.Fatalf("popped %d distinct values, want %d", len(seen), n)
	}
	for v, c := range seen {
		if c != 1 {
			t.Fatalf("value %d popped %d times", v, c)
		}
	}
}

func TestConcurrentPushPop(t *testing.T) {
	var s Stack
	const n = 2000

	var wg sync.WaitGroup
	popped := make([]int, n)
	var mu sync.Mutex
	got := make(map[int]int)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Push(i)
		}(i)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				v, ok := s.Pop()
				if ok {
					mu.Lock()
					got[v]++
					mu.Unlock()
					return
				}
			}
		}()
	}
	wg.Wait()
	_ = popped

	if s.Len() != 0 {
		t.Errorf("Len = %d, want 0", s.Len())
	}
	if len(got) != n {
		t.Fatalf("popped %d distinct values, want %d", len(got), n)
	}
	for v, c := range got {
		if c != 1 {
			t.Fatalf("value %d popped %d times", v, c)
		}
	}
}

func TestIsPusher(t *testing.T) {
	var s Stack
	var p Pusher = &s
	p.Push(7)
	if v, _ := s.Pop(); v != 7 {
		t.Errorf("Pop = %d, want 7", v)
	}
}

func BenchmarkPushPop(b *testing.B) {
	var s Stack
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Push(1)
			s.Pop()
		}
	})
}
