package appendzero

import (
	"errors"
	"reflect"
	"testing"
)

func TestAppendZeroInts(t *testing.T) {
	s := []int{1}
	if err := AppendZero(&s, 2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []int{1, 0, 0}) {
		t.Errorf("s = %v, want [1 0 0]", s)
	}
}

func TestAppendZeroStrings(t *testing.T) {
	s := []string{"a"}
	if err := AppendZero(&s, 1); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []string{"a", ""}) {
		t.Errorf("s = %q, want [a \"\"]", s)
	}
}

func TestAppendZeroStructs(t *testing.T) {
	type item struct{ A, B int }
	s := []item{{1, 2}}
	if err := AppendZero(&s, 1); err != nil {
		t.Fatal(err)
	}
	if len(s) != 2 || s[1] != (item{}) {
		t.Errorf("s = %v, want a zero item appended", s)
	}
}

func TestAppendZeroNilSlice(t *testing.T) {
	var s []int
	if err := AppendZero(&s, 3); err != nil {
		t.Fatal(err)
	}
	if len(s) != 3 {
		t.Errorf("len = %d, want 3", len(s))
	}
}

func TestAppendZeroNoop(t *testing.T) {
	s := []int{1}
	if err := AppendZero(&s, 0); err != nil || len(s) != 1 {
		t.Errorf("s = %v, err = %v, want [1], nil", s, err)
	}
}

func TestAppendZeroBadTarget(t *testing.T) {
	s := []int{1}
	for _, c := range []any{s, nil, (*[]int)(nil), new(int)} {
		if err := AppendZero(c, 1); !errors.Is(err, ErrTarget) {
			t.Errorf("AppendZero(%#v) = %v, want ErrTarget", c, err)
		}
	}
	if err := AppendZero(&s, -1); !errors.Is(err, ErrTarget) {
		t.Errorf("negative n = %v, want ErrTarget", err)
	}
}
