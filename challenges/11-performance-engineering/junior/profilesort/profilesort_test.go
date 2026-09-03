package profilesort

import (
	"reflect"
	"testing"
)

func TestSortByCum(t *testing.T) {
	in := []Entry{{"a", 1, 5}, {"b", 2, 9}, {"c", 0, 7}}
	got := SortByCum(in)
	want := []Entry{{"b", 2, 9}, {"c", 0, 7}, {"a", 1, 5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortByCum = %v, want %v", got, want)
	}
}

func TestSortByCumTieOnFlatThenName(t *testing.T) {
	in := []Entry{{"z", 1, 5}, {"a", 1, 5}, {"m", 4, 5}}
	got := SortByCum(in)
	want := []Entry{{"m", 4, 5}, {"a", 1, 5}, {"z", 1, 5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortByCum = %v, want %v", got, want)
	}
}

func TestSortByCumDoesNotModifyInput(t *testing.T) {
	in := []Entry{{"a", 1, 5}, {"b", 2, 9}}
	before := append([]Entry(nil), in...)
	SortByCum(in)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input was modified: %v, want %v", in, before)
	}
}

func TestSortByCumEmpty(t *testing.T) {
	got := SortByCum(nil)
	if got == nil || len(got) != 0 {
		t.Errorf("SortByCum(nil) = %v, want empty non-nil slice", got)
	}
}
