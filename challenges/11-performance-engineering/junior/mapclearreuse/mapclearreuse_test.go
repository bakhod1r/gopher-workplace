package mapclearreuse

import (
	"reflect"
	"testing"
)

var sink map[string]int

func TestResetEmptiesInPlace(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	alias := m
	Reset(m)
	if len(m) != 0 {
		t.Errorf("m = %v, want empty", m)
	}
	if len(alias) != 0 {
		t.Errorf("the caller's other reference still sees %v — Reset must empty the map itself", alias)
	}
}

func TestResetNilIsSafe(t *testing.T) {
	var m map[string]int
	Reset(m)
	if m != nil {
		t.Errorf("m = %v, want nil", m)
	}
}

func TestTally(t *testing.T) {
	m := map[string]int{"stale": 7}
	got := Tally(m, []string{"a", "a", "b"})
	want := map[string]int{"a": 2, "b": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Tally = %v, want %v (previous contents must be gone)", got, want)
	}
}

func TestTallyReusesTheMap(t *testing.T) {
	m := make(map[string]int, 1000)
	words := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		words = append(words, string(rune('a'+i%26))+string(rune('a'+i/26)))
	}
	allocs := testing.AllocsPerRun(20, func() { sink = Tally(m, words) })
	if allocs > 2 {
		t.Errorf("Tally made %v allocations, want at most 2 — the map must be reused, not recreated", allocs)
	}
}
