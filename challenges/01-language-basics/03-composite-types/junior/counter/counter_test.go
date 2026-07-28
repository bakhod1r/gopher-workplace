package counter

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count([]string{"a", "b", "a", "c", "a", "b"})
	want := map[string]int{"a": 3, "b": 2, "c": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Count=%v; want %v", got, want)
	}
	if len(Count(nil)) != 0 {
		t.Error("nil -> empty")
	}
}
