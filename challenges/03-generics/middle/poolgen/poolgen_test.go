package poolgen

import "testing"

func TestPoolReuses(t *testing.T) {
	built := 0
	p := NewPool(func() []byte { built++; return make([]byte, 4) })
	b := p.Get()
	if built != 1 {
		t.Fatalf("built = %d, want 1", built)
	}
	p.Put(b)
	if p.Idle() != 1 {
		t.Errorf("Idle() = %d, want 1", p.Idle())
	}
	p.Get()
	if built != 1 {
		t.Errorf("built = %d, want 1 (the pooled value should be reused)", built)
	}
	if p.Idle() != 0 {
		t.Errorf("Idle() = %d, want 0", p.Idle())
	}
}

func TestPoolBuildsWhenEmpty(t *testing.T) {
	built := 0
	p := NewPool(func() int { built++; return 7 })
	if got := p.Get(); got != 7 {
		t.Errorf("Get() = %v, want 7", got)
	}
	p.Get()
	if built != 2 {
		t.Errorf("built = %d, want 2", built)
	}
}
