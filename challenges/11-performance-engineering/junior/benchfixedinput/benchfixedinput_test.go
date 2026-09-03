package benchfixedinput

import (
	"reflect"
	"testing"
)

func want(seed uint32, n int) []uint32 {
	out := make([]uint32, 0, n)
	s := seed
	for i := 0; i < n; i++ {
		s = s*1664525 + 1013904223
		out = append(out, s)
	}
	return out
}

func TestFixedInputIsDeterministic(t *testing.T) {
	a := FixedInput(42, 8)
	b := FixedInput(42, 8)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("same seed gave different data:\n%v\n%v", a, b)
	}
}

func TestFixedInputSequence(t *testing.T) {
	if got := FixedInput(1, 3); !reflect.DeepEqual(got, want(1, 3)) {
		t.Errorf("FixedInput(1, 3) = %v, want %v", got, want(1, 3))
	}
	if got := FixedInput(42, 5); !reflect.DeepEqual(got, want(42, 5)) {
		t.Errorf("FixedInput(42, 5) = %v, want %v", got, want(42, 5))
	}
}

func TestFixedInputSeedMatters(t *testing.T) {
	if reflect.DeepEqual(FixedInput(1, 4), FixedInput(2, 4)) {
		t.Error("different seeds produced identical data")
	}
}

func TestFixedInputEmpty(t *testing.T) {
	got := FixedInput(1, 0)
	if got == nil || len(got) != 0 {
		t.Errorf("FixedInput(1, 0) = %v, want empty non-nil slice", got)
	}
}
