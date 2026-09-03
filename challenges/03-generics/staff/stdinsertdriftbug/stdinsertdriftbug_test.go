package stdinsertdriftbug

import (
	"reflect"
	"testing"
	"time"
)

func TestInsertMarksTwo(t *testing.T) {
	got := InsertMarks([]int{1, 2, 3, 4}, []int{1, 3}, 0)
	if !reflect.DeepEqual(got, []int{1, 0, 2, 3, 0, 4}) {
		t.Errorf("InsertMarks = %v, want [1 0 2 3 0 4]", got)
	}
}

func TestInsertMarksEnds(t *testing.T) {
	got := InsertMarks([]int{1, 2}, []int{0, 2}, 9)
	if !reflect.DeepEqual(got, []int{9, 1, 2, 9}) {
		t.Errorf("InsertMarks = %v, want [9 1 2 9]", got)
	}
}

func TestInsertMarksSkipsOutOfRange(t *testing.T) {
	in := []int{1, 2}
	got := InsertMarks(in, []int{5, -1}, 9)
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("InsertMarks = %v, want [1 2]", got)
	}
	if !reflect.DeepEqual(in, []int{1, 2}) {
		t.Errorf("input mutated: %v", in)
	}
}

func TestInsertMarksAtScale(t *testing.T) {
	const n = 100000
	const marks = 500
	in := make([]int, n)
	for i := range in {
		in[i] = i
	}
	at := make([]int, 0, marks)
	for i := 0; i < marks; i++ {
		at = append(at, i*(n/marks))
	}
	start := time.Now()
	got := InsertMarks(in, at, -1)
	if len(got) != n+marks {
		t.Fatalf("length = %d, want %d", len(got), n+marks)
	}
	for i, p := range at {
		if got[p+i] != -1 {
			t.Fatalf("mark %d missing at index %d (got %d)", i, p+i, got[p+i])
		}
	}
	if d := time.Since(start); d > 3*time.Second {
		t.Fatalf("scale insert took %v, want under 3s", d)
	}
}
