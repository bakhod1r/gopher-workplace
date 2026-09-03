package poolwrapper

import (
	"sync"
	"testing"
)

func TestGetReturnsEmptyBuffer(t *testing.T) {
	var p Pool
	b := p.Get()
	if len(b) != 0 {
		t.Errorf("len = %d, want 0", len(b))
	}
	if cap(b) < 1024 {
		t.Errorf("cap = %d, want at least the default 1024", cap(b))
	}
}

func TestGetHonoursSize(t *testing.T) {
	p := Pool{Size: 4096}
	if got := cap(p.Get()); got < 4096 {
		t.Errorf("cap = %d, want at least 4096", got)
	}
}

func TestRecycledBufferComesBackEmpty(t *testing.T) {
	var p Pool
	b := p.Get()
	b = append(b, "secret"...)
	p.Put(b)
	again := p.Get()
	if len(again) != 0 {
		t.Errorf("recycled buffer has len %d and content %q, want an empty buffer", len(again), again)
	}
}

func TestPutNilIsSafe(t *testing.T) {
	var p Pool
	p.Put(nil)
	if got := p.Get(); len(got) != 0 {
		t.Errorf("Get after Put(nil) = %v, want an empty buffer", got)
	}
}

func TestPoolIsConcurrencySafe(t *testing.T) {
	var p Pool
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				b := p.Get()
				b = append(b, byte(j))
				if len(b) != 1 {
					t.Errorf("buffer came back dirty: len %d", len(b))
					return
				}
				p.Put(b)
			}
		}()
	}
	wg.Wait()
}
