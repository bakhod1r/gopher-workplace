package buildergen

import (
	"reflect"
	"testing"
)

func TestBuilderChaining(t *testing.T) {
	got := (&Builder[int]{}).With(1).With(2).With(3).Build()
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Build() = %v, want [1 2 3]", got)
	}
}

func TestBuilderWithAll(t *testing.T) {
	got := (&Builder[string]{}).WithAll("a", "b").With("c").Build()
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Build() = %v, want [a b c]", got)
	}
}

func TestBuilderEmpty(t *testing.T) {
	got := (&Builder[int]{}).Build()
	if got == nil || len(got) != 0 {
		t.Errorf("Build() = %v, want an empty non-nil slice", got)
	}
}

func TestBuildIsACopy(t *testing.T) {
	b := &Builder[int]{}
	b.With(1)
	first := b.Build()
	b.With(2)
	if len(first) != 1 || first[0] != 1 {
		t.Errorf("an earlier Build() result changed: %v, want [1]", first)
	}
}
