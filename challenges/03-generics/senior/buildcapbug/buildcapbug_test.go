package buildcapbug

import (
	"reflect"
	"testing"
)

func TestBuildContents(t *testing.T) {
	var b Builder[int]
	got := b.Add(1).Add(2).Build()
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Build = %v, want [1 2]", got)
	}
}

func TestBuildResultIsIndependent(t *testing.T) {
	var b Builder[int]
	b.Add(1).Add(2)
	got := b.Build()
	got[0] = 99
	after := b.Build()
	if !reflect.DeepEqual(after, []int{1, 2}) {
		t.Errorf("Build = %v, want [1 2]; the built slice shared the builder's storage", after)
	}
}

func TestBuildEmpty(t *testing.T) {
	var b Builder[int]
	if got := b.Build(); len(got) != 0 {
		t.Errorf("Build = %v, want []", got)
	}
}
