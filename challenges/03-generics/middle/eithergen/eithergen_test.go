package eithergen

import "testing"

func TestLeft(t *testing.T) {
	l, r, isLeft := Left[string, int]("boom").Unwrap()
	if l != "boom" || r != 0 || !isLeft {
		t.Errorf("Unwrap() = %q, %v, %v, want boom, 0, true", l, r, isLeft)
	}
}

func TestRight(t *testing.T) {
	l, r, isLeft := Right[string, int](5).Unwrap()
	if l != "" || r != 5 || isLeft {
		t.Errorf("Unwrap() = %q, %v, %v, want \"\", 5, false", l, r, isLeft)
	}
}

func TestEitherZeroValue(t *testing.T) {
	var e Either[string, int]
	l, r, isLeft := e.Unwrap()
	if l != "" || r != 0 || isLeft {
		t.Errorf("zero Either = %q, %v, %v, want zero, zero, false", l, r, isLeft)
	}
}
