package slicessortbasic

import (
	"reflect"
	"testing"
)

func TestSortNames(t *testing.T) {
	s := []string{"c", "a", "b"}
	SortNames(s)
	if !reflect.DeepEqual(s, []string{"a", "b", "c"}) {
		t.Errorf("after SortNames = %v, want [a b c] (sort in place)", s)
	}
}

func TestSortNamesEdges(t *testing.T) {
	empty := []string{}
	SortNames(empty)
	if len(empty) != 0 {
		t.Errorf("empty slice changed: %v", empty)
	}
	one := []string{"a"}
	SortNames(one)
	if !reflect.DeepEqual(one, []string{"a"}) {
		t.Errorf("one-element slice = %v, want [a]", one)
	}
}
