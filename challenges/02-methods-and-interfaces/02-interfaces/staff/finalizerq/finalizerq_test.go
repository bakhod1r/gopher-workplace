package finalizerq

import (
	"runtime"
	"testing"
	"time"
)

func forceGC() {
	for i := 0; i < 20; i++ {
		runtime.GC()
		time.Sleep(2 * time.Millisecond)
	}
}

func TestCloseReleasesOnce(t *testing.T) {
	p := &Pool{}
	h := NewHandle(p)

	h.Close()
	h.Close()
	h.Close()

	if got := p.Released(); got != 1 {
		t.Errorf("Released = %d, want 1", got)
	}
	if !h.Closed() {
		t.Error("Closed = false after Close")
	}
	runtime.KeepAlive(h)
}

func TestNotClosedInitially(t *testing.T) {
	p := &Pool{}
	h := NewHandle(p)
	if h.Closed() {
		t.Error("Closed = true on a fresh handle")
	}
	if p.Released() != 0 {
		t.Errorf("Released = %d, want 0", p.Released())
	}
	h.Close()
}

func TestFinalizerReleasesDroppedHandle(t *testing.T) {
	p := &Pool{}

	func() {
		h := NewHandle(p)
		_ = h // dropped without Close
	}()

	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		forceGC()
		if p.Released() >= 1 {
			break
		}
	}

	if got := p.Released(); got != 1 {
		t.Errorf("Released = %d, want 1 (the finalizer should have run)", got)
	}
}

func TestCloseThenGCReleasesOnce(t *testing.T) {
	p := &Pool{}

	func() {
		h := NewHandle(p)
		h.Close()
	}()

	forceGC()

	if got := p.Released(); got != 1 {
		t.Errorf("Released = %d, want exactly 1", got)
	}
}

func TestManyHandles(t *testing.T) {
	p := &Pool{}
	for i := 0; i < 100; i++ {
		h := NewHandle(p)
		h.Close()
	}
	if got := p.Released(); got != 100 {
		t.Errorf("Released = %d, want 100", got)
	}
}

func TestIsResource(t *testing.T) {
	var r Resource = &Pool{}
	r.Release()
	if r.(*Pool).Released() != 1 {
		t.Error("Release did not register")
	}
}
