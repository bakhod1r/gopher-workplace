package callgraphedge

import (
	"reflect"
	"testing"
)

func TestEdges(t *testing.T) {
	got := Edges([]Sample{{[]string{"a", "b", "c"}, 5}})
	want := []Edge{{"a", "b", 5}, {"b", "c", 5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Edges = %v, want %v", got, want)
	}
}

func TestEdgesSumAcrossSamples(t *testing.T) {
	got := Edges([]Sample{
		{[]string{"main", "a"}, 3},
		{[]string{"main", "a"}, 4},
		{[]string{"main", "b"}, 1},
	})
	want := []Edge{{"main", "a", 7}, {"main", "b", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Edges = %v, want %v", got, want)
	}
}

func TestEdgesCountRecursionPerOccurrence(t *testing.T) {
	got := Edges([]Sample{{[]string{"rec", "rec", "rec"}, 2}})
	want := []Edge{{"rec", "rec", 4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Edges = %v, want %v (two adjacent pairs, 2 each)", got, want)
	}
}

func TestEdgesSkipShortAndJunkSamples(t *testing.T) {
	got := Edges([]Sample{
		{[]string{"solo"}, 5},
		{[]string{"a", "b"}, 0},
		{nil, 3},
	})
	if got == nil || len(got) != 0 {
		t.Errorf("Edges = %v, want empty non-nil slice", got)
	}
}

func TestEdgesOrdering(t *testing.T) {
	got := Edges([]Sample{
		{[]string{"z", "a"}, 1},
		{[]string{"a", "z"}, 1},
		{[]string{"m", "m2"}, 9},
	})
	want := []Edge{{"m", "m2", 9}, {"a", "z", 1}, {"z", "a", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Edges = %v, want %v", got, want)
	}
}

func TestCalleesOf(t *testing.T) {
	edges := []Edge{{"a", "b", 5}, {"a", "c", 9}, {"x", "y", 100}}
	if got := CalleesOf(edges, "a"); !reflect.DeepEqual(got, []string{"c", "b"}) {
		t.Errorf("CalleesOf = %v, want [c b]", got)
	}
	if got := CalleesOf(edges, "nope"); got == nil || len(got) != 0 {
		t.Errorf("CalleesOf = %v, want empty non-nil slice", got)
	}
}
