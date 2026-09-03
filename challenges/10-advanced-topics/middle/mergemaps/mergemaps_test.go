package mergemaps

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	dst := map[string]int{"a": 1}
	added := Merge(dst, map[string]int{"a": 9, "b": 2})
	if added != 1 {
		t.Errorf("added = %d, want 1", added)
	}
	want := map[string]int{"a": 9, "b": 2}
	if !reflect.DeepEqual(dst, want) {
		t.Errorf("dst = %v, want %v", dst, want)
	}
}

func TestMergeIsVisibleToTheCaller(t *testing.T) {
	dst := map[string]int{}
	alias := dst
	Merge(dst, map[string]int{"x": 1})
	if alias["x"] != 1 {
		t.Error("the merge was not applied to the caller's map")
	}
}

func TestMergeEmpty(t *testing.T) {
	dst := map[string]int{"a": 1}
	if got := Merge(dst, nil); got != 0 {
		t.Errorf("added = %d, want 0", got)
	}
	if len(dst) != 1 {
		t.Errorf("dst = %v, want it unchanged", dst)
	}
	if got := Merge(nil, map[string]int{"a": 1}); got != 0 {
		t.Errorf("added = %d, want 0 for a nil dst", got)
	}
}

func TestMergeDoesNotTouchSrc(t *testing.T) {
	src := map[string]int{"a": 1}
	Merge(map[string]int{"a": 5}, src)
	if src["a"] != 1 {
		t.Errorf("src[a] = %d, want 1", src["a"])
	}
}

func TestMergeCountsOnlyNewKeys(t *testing.T) {
	dst := map[string]int{"a": 1, "b": 2}
	if got := Merge(dst, map[string]int{"a": 9, "b": 9}); got != 0 {
		t.Errorf("added = %d, want 0", got)
	}
}
