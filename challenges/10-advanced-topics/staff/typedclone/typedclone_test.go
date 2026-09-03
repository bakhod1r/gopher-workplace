package typedclone

import (
	"reflect"
	"testing"
)

type point struct{ X, Y int }

func TestCloneMapContents(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	got := CloneMap(m)
	if !reflect.DeepEqual(got, m) {
		t.Errorf("CloneMap = %v, want %v", got, m)
	}
}

func TestCloneMapIsIndependent(t *testing.T) {
	m := map[string]int{"a": 1}
	got := CloneMap(m)
	got["b"] = 2
	if _, ok := m["b"]; ok {
		t.Error("the clone shares the original map")
	}
	m["c"] = 3
	if _, ok := got["c"]; ok {
		t.Error("the original shares the clone")
	}
}

func TestCloneMapNil(t *testing.T) {
	var m map[string]int
	if got := CloneMap(m); got != nil {
		t.Errorf("CloneMap(nil) = %v, want nil", got)
	}
}

func TestCloneMapEmpty(t *testing.T) {
	got := CloneMap(map[string]int{})
	if got == nil {
		t.Fatal("CloneMap of an empty map returned nil, want an empty map")
	}
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}

func TestCloneMapOtherTypes(t *testing.T) {
	m := map[int]point{1: {X: 1, Y: 2}}
	got := CloneMap(m)
	if got[1] != (point{X: 1, Y: 2}) {
		t.Errorf("got[1] = %v, want {1 2}", got[1])
	}
	got[1] = point{X: 9}
	if m[1].X != 1 {
		t.Error("struct values are shared")
	}
}

func TestCloneMapIsShallow(t *testing.T) {
	m := map[string][]int{"a": {1, 2}}
	got := CloneMap(m)
	got["a"][0] = 99
	if m["a"][0] != 99 {
		t.Error("the slice value was copied; a shallow clone shares it")
	}
}

func TestCloneMapAllocationsAreBounded(t *testing.T) {
	m := make(map[int]int, 64)
	for i := 0; i < 64; i++ {
		m[i] = i
	}
	var sink map[int]int
	if n := testing.AllocsPerRun(50, func() { sink = CloneMap(m) }); n > 6 {
		t.Errorf("CloneMap made %v allocations for 64 entries, want a handful: size the map up front", n)
	}
	_ = sink
}
