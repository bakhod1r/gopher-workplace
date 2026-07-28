package invert

import (
	"reflect"
	"testing"
)

func TestInvert(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	got := Invert(m)
	want := map[int]string{1: "one", 2: "two", 3: "three"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Invert=%v; want %v", got, want)
	}
	if len(Invert(map[string]int{})) != 0 {
		t.Error("empty -> empty")
	}
}
