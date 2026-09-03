package callfn

import (
	"errors"
	"reflect"
	"testing"
)

func TestCallInts(t *testing.T) {
	got, err := CallInts(func(a, b int) int { return a + b }, 1, 2)
	if err != nil || !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("CallInts = %v, %v, want [3], nil", got, err)
	}
}

func TestCallIntsMultipleResults(t *testing.T) {
	got, err := CallInts(func(a int) (int, int) { return a, -a }, 5)
	if err != nil || !reflect.DeepEqual(got, []int{5, -5}) {
		t.Errorf("CallInts = %v, %v, want [5 -5], nil", got, err)
	}
}

func TestCallIntsNoArgsNoResults(t *testing.T) {
	got, err := CallInts(func() {})
	if err != nil || len(got) != 0 {
		t.Errorf("CallInts = %v, %v, want empty, nil", got, err)
	}
}

func TestCallIntsBadSignatures(t *testing.T) {
	cases := []struct {
		name string
		fn   any
		args []int
	}{
		{"not a func", 3, nil},
		{"nil", nil, nil},
		{"wrong arity", func(a, b int) int { return a }, []int{1}},
		{"string param", func(s string) int { return 0 }, []int{1}},
		{"string result", func(a int) string { return "" }, []int{1}},
		{"variadic", func(a ...int) int { return 0 }, []int{1}},
	}
	for _, c := range cases {
		if _, err := CallInts(c.fn, c.args...); !errors.Is(err, ErrSignature) {
			t.Errorf("%s: err = %v, want ErrSignature", c.name, err)
		}
	}
}
