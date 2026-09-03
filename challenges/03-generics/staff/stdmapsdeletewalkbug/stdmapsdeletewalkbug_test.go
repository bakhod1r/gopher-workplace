package stdmapsdeletewalkbug

import (
	"reflect"
	"testing"
)

func TestPruneReportsEveryRemovedKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 5, "c": 2}
	got := Prune(m, 3)
	if !reflect.DeepEqual(got, []string{"a", "c"}) {
		t.Errorf("Prune = %v, want [a c]", got)
	}
	if !reflect.DeepEqual(m, map[string]int{"b": 5}) {
		t.Errorf("map = %v, want map[b:5]", m)
	}
}

func TestPruneKeepsSurvivors(t *testing.T) {
	m := map[string]int{"a": 5}
	got := Prune(m, 3)
	if len(got) != 0 {
		t.Errorf("Prune = %v, want []", got)
	}
	if len(m) != 1 {
		t.Errorf("map = %v, want map[a:5]", m)
	}
}

func TestPruneEmpty(t *testing.T) {
	if got := Prune(map[string]int{}, 3); len(got) != 0 {
		t.Errorf("Prune = %v, want []", got)
	}
}
