package selfvscum

import (
	"reflect"
	"testing"
)

func TestAnalyze(t *testing.T) {
	got := Analyze([]Sample{{[]string{"main", "a"}, 5}})
	want := map[string]Node{"main": {Flat: 0, Cum: 5}, "a": {Flat: 5, Cum: 5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Analyze = %v, want %v", got, want)
	}
}

func TestAnalyzeAccumulates(t *testing.T) {
	got := Analyze([]Sample{
		{[]string{"main", "a"}, 3},
		{[]string{"main", "b", "a"}, 4},
		{[]string{"main", "b"}, 2},
	})
	want := map[string]Node{
		"main": {Flat: 0, Cum: 9},
		"a":    {Flat: 7, Cum: 7},
		"b":    {Flat: 2, Cum: 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Analyze = %v, want %v", got, want)
	}
}

func TestAnalyzeRecursionCountsCumOnce(t *testing.T) {
	got := Analyze([]Sample{{[]string{"rec", "rec", "rec"}, 6}})
	want := map[string]Node{"rec": {Flat: 6, Cum: 6}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Analyze = %v, want %v", got, want)
	}
}

func TestAnalyzeSkipsJunk(t *testing.T) {
	got := Analyze([]Sample{{[]string{"a"}, 0}, {nil, 3}})
	if got == nil || len(got) != 0 {
		t.Errorf("Analyze = %v, want empty non-nil map", got)
	}
}

func TestLeaves(t *testing.T) {
	nodes := map[string]Node{
		"a": {Flat: 5, Cum: 5},
		"b": {Flat: 0, Cum: 9},
		"c": {Flat: 5, Cum: 5},
		"d": {Flat: 1, Cum: 1},
	}
	if got := Leaves(nodes, 1); !reflect.DeepEqual(got, []string{"a", "c", "d"}) {
		t.Errorf("Leaves = %v, want [a c d]", got)
	}
	if got := Leaves(nodes, 6); got == nil || len(got) != 0 {
		t.Errorf("Leaves = %v, want empty non-nil slice", got)
	}
}
