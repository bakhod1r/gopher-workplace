package profilefilter

import (
	"reflect"
	"testing"
)

func TestFilterDropsSmallRows(t *testing.T) {
	in := []Entry{{"a", 50}, {"b", 1}, {"c", 20}}
	got := Filter(in, 100, 5)
	want := []Entry{{"a", 50}, {"c", 20}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter = %v, want %v", got, want)
	}
}

func TestFilterKeepsExactThreshold(t *testing.T) {
	in := []Entry{{"a", 5}}
	got := Filter(in, 100, 5)
	if !reflect.DeepEqual(got, in) {
		t.Errorf("Filter = %v, want %v (exactly at the threshold is kept)", got, in)
	}
}

func TestFilterPreservesOrder(t *testing.T) {
	in := []Entry{{"c", 30}, {"a", 40}, {"b", 30}}
	got := Filter(in, 100, 10)
	if !reflect.DeepEqual(got, in) {
		t.Errorf("Filter = %v, want the input order %v", got, in)
	}
}

func TestFilterDoesNotModifyInput(t *testing.T) {
	in := []Entry{{"a", 50}, {"b", 1}}
	before := append([]Entry(nil), in...)
	Filter(in, 100, 5)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input was modified: %v, want %v", in, before)
	}
}

func TestFilterNonPositiveTotal(t *testing.T) {
	got := Filter([]Entry{{"a", 50}}, 0, 5)
	if got == nil || len(got) != 0 {
		t.Errorf("Filter(_, 0, _) = %v, want empty non-nil slice", got)
	}
}
