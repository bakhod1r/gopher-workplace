package minbygen

import "testing"

type worker struct {
	name string
	load int
}

func loadOf(w worker) int { return w.load }

func TestMinBy(t *testing.T) {
	ws := []worker{{"a", 5}, {"b", 2}, {"c", 9}}
	got, ok := MinBy(ws, loadOf)
	if !ok || got.name != "b" {
		t.Errorf("MinBy = %+v, %v, want {b 2}, true", got, ok)
	}
}

func TestMinByTieKeepsEarlier(t *testing.T) {
	ws := []worker{{"a", 2}, {"b", 2}}
	got, _ := MinBy(ws, loadOf)
	if got.name != "a" {
		t.Errorf("MinBy = %+v, want {a 2} (earlier wins ties)", got)
	}
}

func TestMinByEmpty(t *testing.T) {
	got, ok := MinBy([]worker{}, loadOf)
	if ok || got.name != "" {
		t.Errorf("MinBy(empty) = %+v, %v, want zero, false", got, ok)
	}
}

func TestMinByCallsKeyOncePerElement(t *testing.T) {
	calls := 0
	ws := []worker{{"a", 3}, {"b", 1}, {"c", 2}}
	MinBy(ws, func(w worker) int { calls++; return w.load })
	if calls > len(ws) {
		t.Errorf("key called %d times, want at most %d", calls, len(ws))
	}
}
