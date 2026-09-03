package valuevspointer

import "testing"

var sink int

func TestIncMutatesTheReceiver(t *testing.T) {
	var c Counter
	c.Inc()
	c.Inc()
	if got := c.Value(); got != 2 {
		t.Errorf("Value = %d, want 2", got)
	}
}

func TestIncCopyDoesNotMutateTheReceiver(t *testing.T) {
	var c Counter
	got := c.IncCopy()
	if c.Value() != 0 {
		t.Errorf("the original changed: %d, want 0", c.Value())
	}
	if got.Value() != 1 {
		t.Errorf("the returned copy = %d, want 1", got.Value())
	}
}

func TestIncCopyChains(t *testing.T) {
	var c Counter
	c = c.IncCopy().IncCopy().IncCopy()
	if got := c.Value(); got != 3 {
		t.Errorf("Value = %d, want 3", got)
	}
}

func TestPointerMethodsDoNotAllocate(t *testing.T) {
	var c Counter
	if allocs := testing.AllocsPerRun(100, func() { c.Inc() }); allocs != 0 {
		t.Errorf("Inc made %v allocations, want 0", allocs)
	}
	if allocs := testing.AllocsPerRun(100, func() { sink = c.Value() }); allocs != 0 {
		t.Errorf("Value made %v allocations, want 0", allocs)
	}
}
