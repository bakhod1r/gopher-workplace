package selftimeleakbug

import (
	"reflect"
	"testing"
)

func TestSelfTimeCreditsTheLeaf(t *testing.T) {
	got := SelfTime([]Sample{{[]string{"main", "a", "b"}, 5}})
	want := map[string]int64{"b": 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SelfTime = %v, want %v — the leaf is the frame on the CPU", got, want)
	}
}

func TestSelfTimeAccumulatesAcrossStacks(t *testing.T) {
	got := SelfTime([]Sample{
		{[]string{"main", "a"}, 3},
		{[]string{"main", "b", "a"}, 4},
		{[]string{"main", "b"}, 1},
	})
	want := map[string]int64{"a": 7, "b": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SelfTime = %v, want %v", got, want)
	}
}

func TestSelfTimeSingleFrame(t *testing.T) {
	got := SelfTime([]Sample{{[]string{"main"}, 2}})
	want := map[string]int64{"main": 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SelfTime = %v, want %v", got, want)
	}
}

func TestSelfTimeDoesNotBlameTheEntryPoint(t *testing.T) {
	got := SelfTime([]Sample{
		{[]string{"main", "handler", "parse"}, 10},
		{[]string{"main", "handler", "encode"}, 20},
	})
	if _, ok := got["main"]; ok {
		t.Errorf("main has self time %d; it never ran on the CPU in these samples", got["main"])
	}
	if got["parse"] != 10 || got["encode"] != 20 {
		t.Errorf("SelfTime = %v, want parse 10 and encode 20", got)
	}
}

func TestSelfTimeSkipsJunk(t *testing.T) {
	got := SelfTime([]Sample{{[]string{"a"}, 0}, {nil, 5}})
	if got == nil || len(got) != 0 {
		t.Errorf("SelfTime = %v, want empty non-nil map", got)
	}
}
