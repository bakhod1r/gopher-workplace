package nilmap

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count([]int{1, 2, 2, 3})
	want := map[int]int{1: 1, 2: 2, 3: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Count=%v; want %v", got, want)
	}
	if got := Count(nil); len(got) != 0 {
		t.Errorf("empty=%v", got)
	}
}
