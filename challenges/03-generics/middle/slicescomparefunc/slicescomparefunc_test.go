package slicescomparefunc

import "testing"

func TestCompareNames(t *testing.T) {
	a := []Item{{"a", 1}}
	b := []Item{{"b", 1}}
	if got := CompareNames(a, b); got >= 0 {
		t.Errorf("CompareNames(a, b) = %d, want negative", got)
	}
	if got := CompareNames(b, a); got <= 0 {
		t.Errorf("CompareNames(b, a) = %d, want positive", got)
	}
	if got := CompareNames(a, a); got != 0 {
		t.Errorf("CompareNames(a, a) = %d, want 0", got)
	}
}

func TestCompareNamesPrefix(t *testing.T) {
	a := []Item{{"a", 1}}
	ab := []Item{{"a", 1}, {"b", 2}}
	if got := CompareNames(a, ab); got >= 0 {
		t.Errorf("CompareNames(prefix, longer) = %d, want negative", got)
	}
	if got := CompareNames(ab, a); got <= 0 {
		t.Errorf("CompareNames(longer, prefix) = %d, want positive", got)
	}
	if got := CompareNames(nil, []Item{}); got != 0 {
		t.Errorf("CompareNames(nil, []) = %d, want 0", got)
	}
}
