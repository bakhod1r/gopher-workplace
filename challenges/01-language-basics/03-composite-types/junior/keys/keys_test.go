package keys

import (
	"reflect"
	"testing"
)

func TestSorted(t *testing.T) {
	m := map[string]int{"banana": 1, "apple": 2, "cherry": 3}
	got := Sorted(m)
	want := []string{"apple", "banana", "cherry"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sorted=%v; want %v", got, want)
	}
	if len(Sorted(map[string]int{})) != 0 {
		t.Error("empty map -> empty keys")
	}
}
