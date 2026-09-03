package closureretain

import (
	"runtime"
	"testing"
)

func TestSummarize(t *testing.T) {
	batch := []Record{{Size: 1}, {Size: 2}, {Size: 3}}
	f := Summarize(batch)
	if got := f(); got != 6 {
		t.Errorf("f() = %d, want 6", got)
	}
	if got := f(); got != 6 {
		t.Errorf("second call = %d, want 6", got)
	}
}

func TestSummarizeEmpty(t *testing.T) {
	if got := Summarize(nil)(); got != 0 {
		t.Errorf("f() = %d, want 0", got)
	}
}

func TestSummarizeSnapshotsTheBatch(t *testing.T) {
	batch := []Record{{Size: 1}, {Size: 2}}
	f := Summarize(batch)
	batch[0].Size = 100
	if got := f(); got != 3 {
		t.Errorf("f() = %d, want 3: the total must be computed before the callback is returned", got)
	}
}

func TestSummarizeReleasesTheBatch(t *testing.T) {
	makeCallback := func() func() int {
		batch := make([]Record, 8192)
		for i := range batch {
			batch[i].Size = 1
		}
		return Summarize(batch)
	}

	runtime.GC()
	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	f := makeCallback()
	if got := f(); got != 8192 {
		t.Fatalf("f() = %d, want 8192", got)
	}

	runtime.GC()
	var after runtime.MemStats
	runtime.ReadMemStats(&after)
	runtime.KeepAlive(f)

	if after.HeapAlloc > before.HeapAlloc+(1<<20) {
		t.Errorf("heap grew by %d bytes with the callback alive, want under 1 MiB: the closure retains the batch",
			after.HeapAlloc-before.HeapAlloc)
	}
}
