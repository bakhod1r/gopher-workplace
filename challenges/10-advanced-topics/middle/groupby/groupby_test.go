package groupby

import (
	"reflect"
	"testing"
)

func TestGroup(t *testing.T) {
	got := Group([][2]int{{1, 10}, {2, 20}, {1, 11}})
	want := map[int][]int{1: {10, 11}, 2: {20}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Group = %v, want %v", got, want)
	}
}

func TestGroupEmpty(t *testing.T) {
	got := Group(nil)
	if got == nil {
		t.Fatal("Group(nil) = nil, want an empty map")
	}
	if len(got) != 0 {
		t.Errorf("Group = %v, want empty", got)
	}
}

func TestGroupPreservesOrder(t *testing.T) {
	pairs := make([][2]int, 0, 100)
	for i := 0; i < 100; i++ {
		pairs = append(pairs, [2]int{i % 3, i})
	}
	got := Group(pairs)
	for k, vs := range got {
		for i := 1; i < len(vs); i++ {
			if vs[i] <= vs[i-1] {
				t.Fatalf("bucket %d is out of order: %v", k, vs)
			}
		}
	}
}

func TestGroupBucketsAreRightSized(t *testing.T) {
	pairs := make([][2]int, 0, 300)
	for i := 0; i < 300; i++ {
		pairs = append(pairs, [2]int{i % 5, i})
	}
	got := Group(pairs)
	for k, vs := range got {
		if cap(vs) != len(vs) {
			t.Errorf("bucket %d has len %d and cap %d: size the buckets from the counts", k, len(vs), cap(vs))
		}
	}
}
