package appendall

import (
	"reflect"
	"testing"
)

var sink []int

func TestAppendAll(t *testing.T) {
	got := AppendAll([][]int{{1}, {2, 3}, {}})
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("AppendAll = %v, want [1 2 3]", got)
	}
	if got := AppendAll(nil); len(got) != 0 {
		t.Errorf("AppendAll = %v, want empty", got)
	}
	if got := AppendAll([][]int{{}, {}}); len(got) != 0 {
		t.Errorf("AppendAll = %v, want empty", got)
	}
}

func TestAppendAllIsIndependent(t *testing.T) {
	a := []int{1, 2}
	got := AppendAll([][]int{a})
	got[0] = 99
	if a[0] != 1 {
		t.Error("the result shares storage with a part")
	}
}

func TestAppendAllAllocatesOnce(t *testing.T) {
	parts := make([][]int, 64)
	for i := range parts {
		parts[i] = []int{i, i, i, i}
	}
	n := testing.AllocsPerRun(50, func() { sink = AppendAll(parts) })
	if n > 1 {
		t.Errorf("AppendAll made %v allocations, want 1: size the result first", n)
	}
}

func TestAppendAllLarge(t *testing.T) {
	parts := make([][]int, 100)
	want := 0
	for i := range parts {
		parts[i] = make([]int, i)
		want += i
	}
	if got := AppendAll(parts); len(got) != want {
		t.Errorf("len = %d, want %d", len(got), want)
	}
}
