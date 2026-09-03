package settableslice

import (
	"errors"
	"reflect"
	"testing"
)

func TestDouble(t *testing.T) {
	s := []int{1, 2, 3}
	if err := Double(s); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []int{2, 4, 6}) {
		t.Errorf("s = %v, want [2 4 6]: the elements were not written", s)
	}
}

func TestDoubleEmptyAndNil(t *testing.T) {
	if err := Double([]int{}); err != nil {
		t.Errorf("Double([]) = %v, want nil", err)
	}
	if err := Double([]int(nil)); err != nil {
		t.Errorf("Double(nil slice) = %v, want nil", err)
	}
}

func TestDoubleWritesThroughAView(t *testing.T) {
	s := []int{1, 2, 3, 4}
	if err := Double(s[1:3]); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []int{1, 4, 6, 4}) {
		t.Errorf("s = %v, want [1 4 6 4]", s)
	}
}

func TestDoubleBadShape(t *testing.T) {
	for _, in := range []any{nil, 3, []string{"a"}, map[string]int{}} {
		if err := Double(in); !errors.Is(err, ErrShape) {
			t.Errorf("Double(%#v) = %v, want ErrShape", in, err)
		}
	}
}
