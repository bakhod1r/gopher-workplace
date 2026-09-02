package bfsgen

import (
	"reflect"
	"testing"
)

func TestBFSOrder(t *testing.T) {
	edges := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"d"},
	}
	got := BFS(edges, "a")
	if !reflect.DeepEqual(got, []string{"a", "b", "c", "d"}) {
		t.Errorf("BFS = %v, want [a b c d]", got)
	}
}

func TestBFSCycles(t *testing.T) {
	edges := map[string][]string{
		"a": {"b"},
		"b": {"a"},
	}
	got := BFS(edges, "a")
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("BFS = %v, want [a b] (each node once)", got)
	}
	self := BFS(map[string][]string{"a": {"a"}}, "a")
	if !reflect.DeepEqual(self, []string{"a"}) {
		t.Errorf("BFS = %v, want [a]", self)
	}
}

func TestBFSIsolatedStart(t *testing.T) {
	got := BFS(map[int][]int{}, 7)
	if !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("BFS = %v, want [7]", got)
	}
}
