package stripedbuffers

import (
	"sync"
	"testing"
	"unsafe"
)

func TestWithRunsOnTheBuffer(t *testing.T) {
	s := NewStriped(4, 64)
	got := s.With(0, func(b []byte) []byte { return append(b, 'a', 'b') })
	if got != 2 {
		t.Errorf("With = %d, want 2", got)
	}
}

func TestBufferIsResetEachTime(t *testing.T) {
	s := NewStriped(2, 64)
	for i := 0; i < 10; i++ {
		got := s.With(0, func(b []byte) []byte { return append(b, 'x') })
		if got != 1 {
			t.Fatalf("call %d: With = %d, want 1: the buffer was not reset", i, got)
		}
	}
}

func TestShardsAreIndependent(t *testing.T) {
	s := NewStriped(4, 64)
	s.With(0, func(b []byte) []byte { return append(b, 'a', 'a', 'a') })
	got := s.With(1, func(b []byte) []byte {
		if len(b) != 0 {
			t.Errorf("shard 1 saw %d bytes from another shard", len(b))
		}
		return append(b, 'b')
	})
	if got != 1 {
		t.Errorf("With = %d, want 1", got)
	}
}

func TestNegativeAndLargeIDs(t *testing.T) {
	s := NewStriped(4, 64)
	for _, id := range []int{-1, -7, 0, 3, 4, 1 << 20} {
		got := s.With(id, func(b []byte) []byte { return append(b, 'z') })
		if got != 1 {
			t.Fatalf("id %d: With = %d, want 1", id, got)
		}
	}
}

func TestStripesDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(stripe{}); got != lineSize {
		t.Errorf("sizeof(stripe) = %d, want %d", got, lineSize)
	}
	s := make([]stripe, 2)
	a := uintptr(unsafe.Pointer(&s[0]))
	b := uintptr(unsafe.Pointer(&s[1]))
	if b-a != lineSize {
		t.Errorf("stride = %d, want %d", b-a, lineSize)
	}
}

func TestConcurrentShards(t *testing.T) {
	s := NewStriped(8, 128)
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				n := s.With(w, func(b []byte) []byte {
					for j := 0; j < 4; j++ {
						b = append(b, byte(w))
					}
					return b
				})
				if n != 4 {
					panic("shard buffer shared between goroutines")
				}
			}
		}(w)
	}
	wg.Wait()
}
