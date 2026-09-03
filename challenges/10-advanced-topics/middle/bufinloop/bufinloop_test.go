package bufinloop

import (
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	got := Render([][]int{{1, 2}, {3}, {}})
	want := []string{"1,2", "3", ""}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Render = %q, want %q", got, want)
	}
	if got := Render(nil); len(got) != 0 {
		t.Errorf("Render = %q, want empty", got)
	}
}

func TestRenderAllocationsScaleWithRowsNotIterations(t *testing.T) {
	rows := make([][]int, 64)
	for i := range rows {
		rows[i] = []int{i, i + 1, i + 2}
	}
	n := testing.AllocsPerRun(50, func() { _ = Render(rows) })
	// one result slice, one scratch buffer, one string per row
	if n > float64(len(rows))+4 {
		t.Errorf("Render made %v allocations for %d rows, want about %d: hoist the scratch buffer", n, len(rows), len(rows)+2)
	}
}
