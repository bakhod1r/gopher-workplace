package chainmapgen

import "testing"

func TestChainGet(t *testing.T) {
	c := NewChain(map[string]int{"a": 1}, map[string]int{"a": 2, "b": 3})
	if v, ok := c.Get("a"); v != 1 || !ok {
		t.Errorf("Get(a) = %v, %v, want 1, true (first layer wins)", v, ok)
	}
	if v, ok := c.Get("b"); v != 3 || !ok {
		t.Errorf("Get(b) = %v, %v, want 3, true", v, ok)
	}
	if v, ok := c.Get("z"); v != 0 || ok {
		t.Errorf("Get(z) = %v, %v, want 0, false", v, ok)
	}
}

func TestChainStoredZeroWins(t *testing.T) {
	c := NewChain(map[string]int{"a": 0}, map[string]int{"a": 9})
	if v, ok := c.Get("a"); v != 0 || !ok {
		t.Errorf("Get(a) = %v, %v, want 0, true", v, ok)
	}
}

func TestChainFlatten(t *testing.T) {
	high := map[string]int{"a": 1}
	low := map[string]int{"a": 2, "b": 3}
	got := NewChain(high, low).Flatten()
	if got["a"] != 1 || got["b"] != 3 || len(got) != 2 {
		t.Errorf("Flatten = %v, want {a:1 b:3}", got)
	}
	if low["a"] != 2 || len(high) != 1 {
		t.Error("Flatten mutated a layer")
	}
}
