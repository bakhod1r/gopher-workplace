package inferenceexplicit

import (
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	var s []int = Empty[int]()
	if s == nil {
		t.Error("Empty[int]() = nil, want an empty non-nil slice")
	}
	if len(s) != 0 {
		t.Errorf("Empty[int]() = %v, want []", s)
	}
	if got := Empty[string](); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("Empty[string]() = %v, want []", got)
	}
}

func TestZeroOf(t *testing.T) {
	if got := ZeroOf[int](); got != 0 {
		t.Errorf("ZeroOf[int]() = %v, want 0", got)
	}
	if got := ZeroOf[string](); got != "" {
		t.Errorf("ZeroOf[string]() = %q, want an empty string", got)
	}
	if got := ZeroOf[bool](); got != false {
		t.Errorf("ZeroOf[bool]() = %v, want false", got)
	}
}

func TestWrapInfers(t *testing.T) {
	if got := Wrap(5); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("Wrap(5) = %v, want [5]", got)
	}
	if got := Wrap("a"); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Wrap(a) = %v, want [a]", got)
	}
}
