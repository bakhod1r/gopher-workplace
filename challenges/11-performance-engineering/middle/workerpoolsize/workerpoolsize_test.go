package workerpoolsize

import (
	"reflect"
	"sync/atomic"
	"testing"
)

func TestMapPreservesOrder(t *testing.T) {
	got := Map([]int{1, 2, 3, 4, 5}, 3, func(v int) int { return v * 2 })
	want := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map = %v, want %v", got, want)
	}
}

func TestMapEdgeCases(t *testing.T) {
	if got := Map(nil, 4, func(v int) int { return v }); got == nil || len(got) != 0 {
		t.Errorf("Map(nil) = %v, want empty non-nil slice", got)
	}
	if got := Map([]int{1, 2}, 0, func(v int) int { return v + 1 }); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("Map with 0 workers = %v, want [2 3]", got)
	}
	if got := Map([]int{1}, 100, func(v int) int { return v }); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Map with more workers than items = %v, want [1]", got)
	}
}

func TestMapRunsEachItemExactlyOnce(t *testing.T) {
	var calls atomic.Int64
	items := make([]int, 1000)
	for i := range items {
		items[i] = i
	}
	got := Map(items, 8, func(v int) int {
		calls.Add(1)
		return v * v
	})
	if calls.Load() != 1000 {
		t.Errorf("f called %d times, want 1000", calls.Load())
	}
	for i, v := range got {
		if v != i*i {
			t.Fatalf("result[%d] = %d, want %d", i, v, i*i)
		}
	}
}

func TestMapRespectsTheWorkerLimit(t *testing.T) {
	var live, peak atomic.Int64
	items := make([]int, 200)
	Map(items, 4, func(v int) int {
		n := live.Add(1)
		for {
			p := peak.Load()
			if n <= p || peak.CompareAndSwap(p, n) {
				break
			}
		}
		live.Add(-1)
		return v
	})
	if peak.Load() > 4 {
		t.Errorf("peak concurrency = %d, want at most 4", peak.Load())
	}
}

func TestSizing(t *testing.T) {
	cases := []struct {
		cpus  int
		block float64
		want  int
	}{
		{8, 0, 8},
		{8, 0.5, 16},
		{8, 0.9, 80},
		{8, -1, 8},
		{8, 1.5, 8},
		{0, 0.5, 1},
		{-4, 0, 1},
	}
	for _, c := range cases {
		if got := Sizing(c.cpus, c.block); got != c.want {
			t.Errorf("Sizing(%d, %v) = %d, want %d", c.cpus, c.block, got, c.want)
		}
	}
}
