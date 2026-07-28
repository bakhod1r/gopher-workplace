package zip

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	got := Zip([]string{"a", "b", "c"}, []int{1, 2, 3})
	want := map[string]int{"a": 1, "b": 2, "c": 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Zip=%v; want %v", got, want)
	}
	got = Zip([]string{"a", "b"}, []int{1, 2, 3})
	if !reflect.DeepEqual(got, map[string]int{"a": 1, "b": 2}) {
		t.Errorf("Zip mismatched lengths: %v", got)
	}
}
