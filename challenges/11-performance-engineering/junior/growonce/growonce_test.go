package growonce

import (
	"reflect"
	"testing"
)

var sink []int

func TestGrowToKeepsContents(t *testing.T) {
	s := []int{1, 2, 3}
	got := GrowTo(s, 100)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("GrowTo = %v, want [1 2 3]", got)
	}
	if cap(got) < 100 {
		t.Errorf("cap = %d, want at least 100", cap(got))
	}
	if len(got) != 3 {
		t.Errorf("len = %d, want 3", len(got))
	}
}

func TestGrowToNoopWhenBigEnough(t *testing.T) {
	s := make([]int, 3, 50)
	got := GrowTo(s, 10)
	if cap(got) != 50 || len(got) != 3 {
		t.Errorf("len, cap = %d, %d, want 3, 50", len(got), cap(got))
	}
	allocs := testing.AllocsPerRun(50, func() { sink = GrowTo(s, 10) })
	if allocs != 0 {
		t.Errorf("GrowTo on a slice that already fits made %v allocations, want 0", allocs)
	}
}

func TestGrowToAllocatesOnce(t *testing.T) {
	s := make([]int, 500)
	allocs := testing.AllocsPerRun(50, func() { sink = GrowTo(s, 5000) })
	if allocs > 1 {
		t.Errorf("GrowTo made %v allocations, want at most 1", allocs)
	}
}

func TestGrowToDoesNotAliasAfterGrowth(t *testing.T) {
	s := []int{1, 2, 3}
	got := GrowTo(s, 100)
	got[0] = 99
	if s[0] != 1 {
		t.Errorf("growth aliased the original: s = %v", s)
	}
}

func TestGrowToNil(t *testing.T) {
	got := GrowTo(nil, 8)
	if len(got) != 0 || cap(got) < 8 {
		t.Errorf("len, cap = %d, %d, want 0 and at least 8", len(got), cap(got))
	}
}
