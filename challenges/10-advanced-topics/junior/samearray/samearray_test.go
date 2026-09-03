package samearray

import "testing"

func TestSameArray(t *testing.T) {
	s := []int{1, 2, 3}
	if !SameArray(s, s) {
		t.Error("SameArray(s, s) = false, want true")
	}
	if !SameArray(s, s[:2]) {
		t.Error("SameArray(s, s[:2]) = false, want true")
	}
	if SameArray(s, s[1:]) {
		t.Error("SameArray(s, s[1:]) = true, want false: the start differs")
	}
}

func TestSameArrayDistinctSlices(t *testing.T) {
	a := []int{1, 2}
	b := []int{1, 2}
	if SameArray(a, b) {
		t.Error("SameArray = true for two separate arrays, want false")
	}
}

func TestSameArrayEmpty(t *testing.T) {
	s := []int{1}
	for _, c := range [][2][]int{
		{nil, nil}, {s, nil}, {nil, s}, {s[:0], s},
	} {
		if SameArray(c[0], c[1]) {
			t.Errorf("SameArray(%v, %v) = true, want false for an empty operand", c[0], c[1])
		}
	}
}

func TestSameArrayAfterCopy(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, 3)
	copy(b, a)
	if SameArray(a, b) {
		t.Error("a copy must not share storage")
	}
}
