package radixlongestbug

import (
	"testing"
)

func TestLongestBeatsCatchAll(t *testing.T) {
	var p Prefixes[int]
	p.Add("", 0)
	p.Add("/a", 1)
	p.Add("/a/b", 2)
	if v, ok := p.Longest("/a/b/c"); !ok || v != 2 {
		t.Errorf("Longest(/a/b/c) = %d, %v, want 2, true", v, ok)
	}
	if v, ok := p.Longest("/a/z"); !ok || v != 1 {
		t.Errorf("Longest(/a/z) = %d, %v, want 1, true", v, ok)
	}
	if v, ok := p.Longest("/z"); !ok || v != 0 {
		t.Errorf("Longest(/z) = %d, %v, want 0, true", v, ok)
	}
}

func TestLongestNoMatch(t *testing.T) {
	var p Prefixes[int]
	p.Add("/a", 1)
	if v, ok := p.Longest("/z"); ok || v != 0 {
		t.Errorf("Longest(/z) = %d, %v, want 0, false", v, ok)
	}
	if p.Len() != 1 {
		t.Errorf("Len = %d, want 1", p.Len())
	}
}

func TestLongestExactAndNested(t *testing.T) {
	var p Prefixes[string]
	p.Add("ab", "short")
	p.Add("abcd", "long")
	if v, ok := p.Longest("abcd"); !ok || v != "long" {
		t.Errorf("Longest(abcd) = %q, %v, want long, true", v, ok)
	}
	if v, ok := p.Longest("abc"); !ok || v != "short" {
		t.Errorf("Longest(abc) = %q, %v, want short, true", v, ok)
	}
	if v, ok := p.Longest("abcdef"); !ok || v != "long" {
		t.Errorf("Longest(abcdef) = %q, %v, want long, true", v, ok)
	}
}
