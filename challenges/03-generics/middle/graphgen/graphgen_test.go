package graphgen

import (
	"reflect"
	"testing"
)

func TestGraphEdges(t *testing.T) {
	g := NewGraph[string]()
	g.AddEdge("a", "b")
	g.AddEdge("a", "c")
	if got := g.Neighbors("a"); !reflect.DeepEqual(got, []string{"b", "c"}) {
		t.Errorf("Neighbors(a) = %v, want [b c]", got)
	}
	if g.Degree("a") != 2 {
		t.Errorf("Degree(a) = %d, want 2", g.Degree("a"))
	}
}

func TestGraphIgnoresDuplicates(t *testing.T) {
	g := NewGraph[string]()
	g.AddEdge("a", "b")
	g.AddEdge("a", "b")
	if g.Degree("a") != 1 {
		t.Errorf("Degree(a) = %d, want 1", g.Degree("a"))
	}
}

func TestGraphUnknownNode(t *testing.T) {
	g := NewGraph[int]()
	if got := g.Neighbors(9); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Neighbors(unknown) = %v, want []", got)
	}
	if g.Degree(9) != 0 {
		t.Errorf("Degree(unknown) = %d, want 0", g.Degree(9))
	}
}

func TestGraphNeighborsIsACopy(t *testing.T) {
	g := NewGraph[string]()
	g.AddEdge("a", "b")
	ns := g.Neighbors("a")
	ns[0] = "zzz"
	if again := g.Neighbors("a"); again[0] != "b" {
		t.Errorf("Neighbors returned internal storage: %v", again)
	}
}
