package slicesclipgen

import (
	"reflect"
	"testing"
)

func TestFreezeCapacity(t *testing.T) {
	buf := make([]int, 8)
	for i := range buf {
		buf[i] = i
	}
	view := buf[0:2]
	if cap(view) == len(view) {
		t.Fatal("test setup: the view should start with spare capacity")
	}
	got := Freeze(view)
	if cap(got) != len(got) {
		t.Errorf("cap = %d, len = %d, want them equal", cap(got), len(got))
	}
	if !reflect.DeepEqual(got, []int{0, 1}) {
		t.Errorf("Freeze = %v, want [0 1]", got)
	}
}

func TestFreezeStopsOverwrites(t *testing.T) {
	buf := make([]int, 4)
	for i := range buf {
		buf[i] = i
	}
	got := append(Freeze(buf[0:2]), 99)
	if buf[2] != 2 {
		t.Errorf("appending overwrote the buffer: buf = %v", buf)
	}
	if len(got) != 3 || got[2] != 99 {
		t.Errorf("append result = %v, want [0 1 99]", got)
	}
}

func TestFreezeNil(t *testing.T) {
	got := Freeze([]string(nil))
	if got == nil || len(got) != 0 {
		t.Errorf("Freeze(nil) = %v, want an empty non-nil slice", got)
	}
}
