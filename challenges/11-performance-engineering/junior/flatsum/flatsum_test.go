package flatsum

import (
	"reflect"
	"testing"
)

func TestFlatSum(t *testing.T) {
	got := FlatSum([]Sample{{"a", 3}, {"b", 1}, {"a", 2}})
	want := map[string]int64{"a": 5, "b": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatSum = %v, want %v", got, want)
	}
}

func TestFlatSumIgnoresNonPositive(t *testing.T) {
	got := FlatSum([]Sample{{"a", 5}, {"b", 0}, {"c", -3}, {"a", -1}})
	want := map[string]int64{"a": 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatSum = %v, want %v", got, want)
	}
}

func TestFlatSumEmpty(t *testing.T) {
	got := FlatSum(nil)
	if got == nil {
		t.Fatal("FlatSum(nil) = nil, want empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("FlatSum(nil) = %v, want empty", got)
	}
}
