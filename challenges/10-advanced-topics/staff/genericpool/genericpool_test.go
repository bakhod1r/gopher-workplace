package genericpool

import (
	"sync"
	"testing"
)

type buffer struct {
	Name  string
	Count int
}

func TestGetReturnsAZeroValue(t *testing.T) {
	p := NewPool[buffer]()
	v := p.Get()
	if v == nil {
		t.Fatal("Get returned nil")
	}
	if *v != (buffer{}) {
		t.Errorf("Get = %+v, want the zero buffer", *v)
	}
}

func TestGetZeroesARecycledValue(t *testing.T) {
	p := NewPool[buffer]()
	v := p.Get()
	v.Name = "dirty"
	v.Count = 9
	p.Put(v)
	got := p.Get()
	if *got != (buffer{}) {
		t.Errorf("Get = %+v, want the zero buffer: a recycled value must be reset", *got)
	}
}

func TestGetIsTypedWithoutAssertions(t *testing.T) {
	p := NewPool[int]()
	v := p.Get()
	*v = 42
	if *v != 42 {
		t.Errorf("*v = %d, want 42", *v)
	}
	p.Put(v)
	if got := p.Get(); *got != 0 {
		t.Errorf("*got = %d, want 0", *got)
	}
}

func TestPutNil(t *testing.T) {
	p := NewPool[buffer]()
	p.Put(nil)
	if v := p.Get(); v == nil {
		t.Error("Get returned nil after Put(nil)")
	}
}

func TestConcurrentGetPut(t *testing.T) {
	p := NewPool[buffer]()
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				v := p.Get()
				if v.Count != 0 || v.Name != "" {
					panic("a dirty value came out of the pool")
				}
				v.Count = w
				v.Name = "in-use"
				if v.Count != w {
					panic("value shared between goroutines")
				}
				p.Put(v)
			}
		}(w)
	}
	wg.Wait()
}
