package builderreuse

import (
	"reflect"
	"testing"
)

func TestRenderLines(t *testing.T) {
	got := RenderLines([][]int{{1, 2}, {3}, {}})
	want := []string{"1-2", "3", ""}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RenderLines = %q, want %q", got, want)
	}
	if got := RenderLines(nil); len(got) != 0 {
		t.Errorf("RenderLines = %q, want empty", got)
	}
}

func TestRenderLinesNegatives(t *testing.T) {
	got := RenderLines([][]int{{-1, 2, -3}})
	if len(got) != 1 || got[0] != "-1-2--3" {
		t.Errorf("RenderLines = %q, want [-1-2--3]", got)
	}
}

func TestRenderLinesDoesNotLeakBetweenRows(t *testing.T) {
	rows := make([][]int, 50)
	for i := range rows {
		rows[i] = []int{i}
	}
	got := RenderLines(rows)
	for i, s := range got {
		want := strconvItoa(i)
		if s != want {
			t.Fatalf("line %d = %q, want %q: the builder was not reset", i, s, want)
		}
	}
}

func strconvItoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

func TestRenderLinesAllocationsScaleWithRows(t *testing.T) {
	rows := make([][]int, 64)
	for i := range rows {
		rows[i] = []int{i, i + 1, i + 2}
	}
	n := testing.AllocsPerRun(50, func() { _ = RenderLines(rows) })
	if n > float64(len(rows))*2+8 {
		t.Errorf("RenderLines made %v allocations for %d rows, want about %d: reuse one builder",
			n, len(rows), len(rows)*2)
	}
}
