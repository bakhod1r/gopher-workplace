package currygen

import (
	"strings"
	"testing"
)

func TestCurry2(t *testing.T) {
	add := func(a, b int) int { return a + b }
	if got := Curry2(add)(1)(2); got != 3 {
		t.Errorf("Curry2(add)(1)(2) = %v, want 3", got)
	}
	if got := Curry2(strings.Repeat)("ab")(2); got != "abab" {
		t.Errorf("Curry2(Repeat)(ab)(2) = %q, want abab", got)
	}
}

func TestCurry2PartialApplication(t *testing.T) {
	add := func(a, b int) int { return a + b }
	plus1 := Curry2(add)(1)
	if got := plus1(5); got != 6 {
		t.Errorf("plus1(5) = %v, want 6", got)
	}
	if got := plus1(0); got != 1 {
		t.Errorf("plus1(0) = %v, want 1", got)
	}
}

func TestCurry2ArgumentOrder(t *testing.T) {
	sub := func(a, b int) int { return a - b }
	if got := Curry2(sub)(10)(3); got != 7 {
		t.Errorf("Curry2(sub)(10)(3) = %v, want 7", got)
	}
}
