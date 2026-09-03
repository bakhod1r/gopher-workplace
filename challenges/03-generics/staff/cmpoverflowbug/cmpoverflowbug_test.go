package cmpoverflowbug

import (
	"math"
	"reflect"
	"testing"
)

type row struct {
	name string
	k    int
}

func names(rs []row) []string {
	out := make([]string, len(rs))
	for i, r := range rs {
		out[i] = r.name
	}
	return out
}

func key(r row) int { return r.k }

func TestSortByKeyOrdinary(t *testing.T) {
	rs := []row{{"c", 3}, {"a", 1}, {"b", 2}}
	SortByKey(rs, key)
	if got := names(rs); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("SortByKey = %v, want [a b c]", got)
	}
}

func TestSortByKeyExtremePair(t *testing.T) {
	rs := []row{{"max", math.MaxInt}, {"min", math.MinInt}}
	SortByKey(rs, key)
	if got := names(rs); !reflect.DeepEqual(got, []string{"min", "max"}) {
		t.Errorf("SortByKey = %v, want [min max]", got)
	}
}

func TestSortByKeyFullRange(t *testing.T) {
	rs := []row{{"zero", 0}, {"max", math.MaxInt}, {"min", math.MinInt}, {"five", 5}, {"negfive", -5}}
	SortByKey(rs, key)
	want := []string{"min", "negfive", "zero", "five", "max"}
	if got := names(rs); !reflect.DeepEqual(got, want) {
		t.Errorf("SortByKey = %v, want %v", got, want)
	}
}
