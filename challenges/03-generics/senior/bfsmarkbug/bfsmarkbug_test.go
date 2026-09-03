package bfsmarkbug

import (
	"reflect"
	"testing"
)

func TestBFSDiamondVisitsOnce(t *testing.T) {
	adj := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"d"},
	}
	got := BFS(adj, "a")
	if !reflect.DeepEqual(got, []string{"a", "b", "c", "d"}) {
		t.Errorf("BFS = %v, want [a b c d]", got)
	}
}

func TestBFSIsolated(t *testing.T) {
	got := BFS(map[string][]string{}, "x")
	if !reflect.DeepEqual(got, []string{"x"}) {
		t.Errorf("BFS = %v, want [x]", got)
	}
}

func TestBFSCycle(t *testing.T) {
	adj := map[int][]int{1: {2}, 2: {1}}
	if got := BFS(adj, 1); len(got) != 2 {
		t.Errorf("BFS = %v, want 2 nodes", got)
	}
}
