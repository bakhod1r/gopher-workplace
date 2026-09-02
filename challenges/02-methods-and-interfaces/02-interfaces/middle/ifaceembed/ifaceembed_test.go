package ifaceembed

import "testing"

func TestGauge(t *testing.T) {
	if got := (Gauge{N: 5}).Value(); got != 5 {
		t.Errorf("Gauge.Value = %d, want 5", got)
	}
}

func TestCountingSource(t *testing.T) {
	c := &CountingSource{Source: Gauge{N: 5}}
	if got := c.Value(); got != 5 {
		t.Errorf("Value = %d, want 5", got)
	}
	if c.Calls != 1 {
		t.Errorf("Calls = %d, want 1", c.Calls)
	}
	c.Value()
	if c.Calls != 2 {
		t.Errorf("Calls = %d, want 2", c.Calls)
	}
}

func TestReadTwice(t *testing.T) {
	c := &CountingSource{Source: Gauge{N: 7}}
	a, b := ReadTwice(c)
	if a != 7 || b != 7 {
		t.Errorf("ReadTwice = %d, %d, want 7, 7", a, b)
	}
	if c.Calls != 2 {
		t.Errorf("Calls = %d, want 2", c.Calls)
	}
}

func TestNested(t *testing.T) {
	inner := &CountingSource{Source: Gauge{N: 1}}
	outer := &CountingSource{Source: inner}
	if got := outer.Value(); got != 1 {
		t.Errorf("Value = %d, want 1", got)
	}
	if outer.Calls != 1 || inner.Calls != 1 {
		t.Errorf("outer=%d inner=%d, want 1 and 1", outer.Calls, inner.Calls)
	}
}
