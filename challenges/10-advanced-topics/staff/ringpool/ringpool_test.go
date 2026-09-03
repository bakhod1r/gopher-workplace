package ringpool

import (
	"sync"
	"testing"
)

func TestGetReturnsASizedBuffer(t *testing.T) {
	p := NewBufPool(4, 64)
	b := p.Get()
	if len(b) != 0 {
		t.Errorf("len = %d, want 0", len(b))
	}
	if cap(b) != 64 {
		t.Errorf("cap = %d, want 64", cap(b))
	}
}

func TestGetRecyclesAPutBuffer(t *testing.T) {
	p := NewBufPool(4, 64)
	b := p.Get()
	b = append(b, 'x')
	p.Put(b)
	got := p.Get()
	if len(got) != 0 {
		t.Errorf("len = %d, want 0: a recycled buffer must be empty", len(got))
	}
	if cap(got) != 64 {
		t.Errorf("cap = %d, want 64", cap(got))
	}
}

func TestGetNeverHandsOutTheSameBufferTwice(t *testing.T) {
	p := NewBufPool(4, 64)
	for i := 0; i < 4; i++ {
		b := make([]byte, 0, 64)
		p.Put(b)
	}
	seen := map[*byte]bool{}
	for i := 0; i < 4; i++ {
		b := p.Get()
		b = append(b, 0)
		if seen[&b[0]] {
			t.Fatal("the same buffer was handed out twice")
		}
		seen[&b[0]] = true
	}
}

func TestGetFallsBackWhenEmpty(t *testing.T) {
	p := NewBufPool(2, 32)
	for i := 0; i < 10; i++ {
		b := p.Get()
		if cap(b) != 32 {
			t.Fatalf("cap = %d, want 32", cap(b))
		}
	}
}

func TestPutRejectsWrongSizes(t *testing.T) {
	p := NewBufPool(2, 32)
	p.Put(make([]byte, 0, 4096))
	b := p.Get()
	if cap(b) != 32 {
		t.Errorf("cap = %d, want 32: an oversized buffer must not enter the ring", cap(b))
	}
}

func TestConcurrentGetAndPut(t *testing.T) {
	p := NewBufPool(8, 128)
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				b := p.Get()
				if cap(b) != 128 || len(b) != 0 {
					panic("bad buffer from the pool")
				}
				b = append(b, byte(w))
				if b[0] != byte(w) {
					panic("buffer shared between goroutines")
				}
				p.Put(b)
			}
		}(w)
	}
	wg.Wait()
}
