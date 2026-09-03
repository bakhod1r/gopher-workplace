package perworkerbuf

import (
	"strconv"
	"strings"
	"testing"
)

func TestRenderAll(t *testing.T) {
	got := RenderAll([][]int{{1, 2}, {3}, {}})
	want := []string{"1,2", "3", ""}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if got := RenderAll(nil); len(got) != 0 {
		t.Errorf("RenderAll = %q, want empty", got)
	}
}

func TestRenderAllUnderConcurrency(t *testing.T) {
	const n = 128
	rows := make([][]int, n)
	for i := range rows {
		rows[i] = []int{i, i * 2, i * 3}
	}
	for round := 0; round < 20; round++ {
		got := RenderAll(rows)
		for i := range rows {
			want := strings.Join([]string{
				strconv.Itoa(i), strconv.Itoa(i * 2), strconv.Itoa(i * 3),
			}, ",")
			if got[i] != want {
				t.Fatalf("round %d: got[%d] = %q, want %q", round, i, got[i], want)
			}
		}
	}
}

func TestRenderAllOrderIsPreserved(t *testing.T) {
	rows := [][]int{{9}, {8}, {7}, {6}}
	got := RenderAll(rows)
	for i, want := range []string{"9", "8", "7", "6"} {
		if got[i] != want {
			t.Fatalf("got[%d] = %q, want %q", i, got[i], want)
		}
	}
}
