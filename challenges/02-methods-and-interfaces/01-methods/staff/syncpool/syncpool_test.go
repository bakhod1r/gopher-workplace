package syncpool

import "testing"

func TestBufferPool(t *testing.T) {
	p := NewPool()
	b := p.Get()
	if len(b.Data) != 1024 {
		t.Fatal("expected 1024")
	}

	p.Put(b)
	b2 := p.Get()
	if b != b2 {
		t.Log("sync.Pool might return a different one, but usually it's the same")
	}
}
