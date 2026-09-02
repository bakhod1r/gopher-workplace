package composegen

import (
	"strconv"
	"testing"
)

func TestCompose(t *testing.T) {
	double := func(n int) int { return n * 2 }
	if got := Compose(double, strconv.Itoa)(2); got != "4" {
		t.Errorf("Compose(double, Itoa)(2) = %q, want \"4\"", got)
	}
	if got := Compose(strconv.Itoa, func(s string) int { return len(s) })(12); got != 2 {
		t.Errorf("Compose(Itoa, len)(12) = %v, want 2", got)
	}
	if got := Compose(double, double)(3); got != 12 {
		t.Errorf("Compose(double, double)(3) = %v, want 12", got)
	}
}

func TestComposeOrder(t *testing.T) {
	var calls []string
	f := func(n int) int { calls = append(calls, "f"); return n }
	g := func(n int) int { calls = append(calls, "g"); return n }
	Compose(f, g)(1)
	if len(calls) != 2 || calls[0] != "f" || calls[1] != "g" {
		t.Errorf("call order = %v, want [f g]", calls)
	}
}
