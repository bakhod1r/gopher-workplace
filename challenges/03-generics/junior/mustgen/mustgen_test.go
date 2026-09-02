package mustgen

import (
	"strings"
	"testing"
)

func TestMustReturnsValue(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}
	if got := Must(Lookup(m, "a")); got != 1 {
		t.Errorf("Must(Lookup(m, a)) = %v, want 1", got)
	}
	if got := Must(Lookup(m, "zero")); got != 0 {
		t.Errorf("Must(Lookup(m, zero)) = %v, want 0", got)
	}
	if got := Must("x", true); got != "x" {
		t.Errorf("Must(x, true) = %q, want x", got)
	}
}

func TestMustPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Must did not panic on ok == false")
		}
		if msg, ok := r.(string); ok && !strings.Contains(msg, "Must") {
			t.Errorf("panic message = %q, want it to mention Must", msg)
		}
	}()
	Must(Lookup(map[string]int{}, "missing"))
}
