package batchflush

import (
	"reflect"
	"testing"
)

func collect() (*[][]int, func([]int)) {
	var got [][]int
	return &got, func(batch []int) {
		got = append(got, append([]int(nil), batch...))
	}
}

func TestAddFlushesWhenFull(t *testing.T) {
	got, flush := collect()
	b := &Batcher{Size: 2, Flush: flush}
	for _, v := range []int{1, 2, 3, 4} {
		b.Add(v)
	}
	want := [][]int{{1, 2}, {3, 4}}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("batches = %v, want %v", *got, want)
	}
	if b.Flushes() != 2 {
		t.Errorf("Flushes = %d, want 2", b.Flushes())
	}
}

func TestCloseFlushesThePartialBatch(t *testing.T) {
	got, flush := collect()
	b := &Batcher{Size: 3, Flush: flush}
	b.Add(1)
	b.Add(2)
	if len(*got) != 0 {
		t.Fatalf("flushed early: %v", *got)
	}
	b.Close()
	if !reflect.DeepEqual(*got, [][]int{{1, 2}}) {
		t.Errorf("batches = %v, want [[1 2]]", *got)
	}
}

func TestCloseIsIdempotentAndNeverFlushesEmpty(t *testing.T) {
	got, flush := collect()
	b := &Batcher{Size: 2, Flush: flush}
	b.Add(1)
	b.Add(2) // full flush
	b.Close()
	b.Close()
	if !reflect.DeepEqual(*got, [][]int{{1, 2}}) {
		t.Errorf("batches = %v, want exactly one batch", *got)
	}
	if b.Flushes() != 1 {
		t.Errorf("Flushes = %d, want 1", b.Flushes())
	}
}

func TestNonPositiveSizeBatchesOne(t *testing.T) {
	got, flush := collect()
	b := &Batcher{Size: 0, Flush: flush}
	b.Add(1)
	b.Add(2)
	b.Close()
	if !reflect.DeepEqual(*got, [][]int{{1}, {2}}) {
		t.Errorf("batches = %v, want [[1] [2]]", *got)
	}
}

func TestNilFlushIsSafe(t *testing.T) {
	b := &Batcher{Size: 2}
	b.Add(1)
	b.Add(2)
	b.Close()
	if b.Flushes() != 1 {
		t.Errorf("Flushes = %d, want 1 — a nil Flush still counts as a flush", b.Flushes())
	}
}

func TestBatcherReusesItsBuffer(t *testing.T) {
	b := &Batcher{Size: 64, Flush: func([]int) {}}
	for i := 0; i < 64; i++ {
		b.Add(i)
	}
	allocs := testing.AllocsPerRun(50, func() {
		for i := 0; i < 64; i++ {
			b.Add(i)
		}
	})
	if allocs != 0 {
		t.Errorf("steady-state batching made %v allocations, want 0", allocs)
	}
}
