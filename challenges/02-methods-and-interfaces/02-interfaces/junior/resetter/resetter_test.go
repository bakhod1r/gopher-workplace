package resetter

import "testing"

func TestReset(t *testing.T) {
	b := &Buffer{Data: []string{"x", "y"}}
	b.Reset()
	if len(b.Data) != 0 {
		t.Errorf("Buffer.Data len = %d, want 0", len(b.Data))
	}

	g := &Gauge{Value: 9}
	g.Reset()
	if g.Value != 0 {
		t.Errorf("Gauge.Value = %d, want 0", g.Value)
	}
}

func TestResetAll(t *testing.T) {
	b := &Buffer{Data: []string{"x"}}
	g := &Gauge{Value: 5}
	ResetAll([]Resetter{b, g})
	if len(b.Data) != 0 || g.Value != 0 {
		t.Errorf("after ResetAll: Data=%v Value=%d", b.Data, g.Value)
	}
	ResetAll(nil)
}

func TestResetIdempotent(t *testing.T) {
	g := &Gauge{}
	g.Reset()
	g.Reset()
	if g.Value != 0 {
		t.Errorf("Value = %d, want 0", g.Value)
	}
}
