package mergesorted

import (
	"reflect"
	"sort"
	"testing"
)

var sink []int

func TestMerge(t *testing.T) {
	got := Merge(nil, [][]int{{1, 3, 5}, {2, 4}, {}})
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Merge = %v, want [1 2 3 4 5]", got)
	}
}

func TestMergeAppendsToDst(t *testing.T) {
	got := Merge([]int{0}, [][]int{{2}, {1}})
	if !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Errorf("Merge = %v, want [0 1 2]", got)
	}
}

func TestMergeEmpty(t *testing.T) {
	if got := Merge(nil, nil); len(got) != 0 {
		t.Errorf("Merge = %v, want empty", got)
	}
	if got := Merge(nil, [][]int{{}, {}}); len(got) != 0 {
		t.Errorf("Merge = %v, want empty", got)
	}
}

func TestMergeDuplicates(t *testing.T) {
	got := Merge(nil, [][]int{{1, 1}, {1}})
	if !reflect.DeepEqual(got, []int{1, 1, 1}) {
		t.Errorf("Merge = %v, want [1 1 1]", got)
	}
}

func TestMergeMatchesSorting(t *testing.T) {
	runs := make([][]int, 7)
	var all []int
	for i := range runs {
		r := make([]int, 0, 20)
		for j := 0; j < 20; j++ {
			r = append(r, i+j*7)
		}
		sort.Ints(r)
		runs[i] = r
		all = append(all, r...)
	}
	sort.Ints(all)
	got := Merge(nil, runs)
	if !reflect.DeepEqual(got, all) {
		t.Errorf("Merge produced a different order than sorting the union")
	}
}

func TestMergeAllocatesNothingWithRoom(t *testing.T) {
	runs := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	dst := make([]int, 0, 32)
	if n := testing.AllocsPerRun(100, func() { sink = Merge(dst[:0], runs) }); n != 0 {
		t.Errorf("Merge made %v allocations, want 0", n)
	}
}
