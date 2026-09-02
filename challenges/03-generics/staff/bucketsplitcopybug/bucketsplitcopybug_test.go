package bucketsplitcopybug

import (
	"testing"
	"time"
)

func TestIndexNoDuplicates(t *testing.T) {
	ix := &Index[int, string]{Max: 4}
	order := []int{7, 3, 15, 1, 20, 9, 12, 4, 18, 6, 11, 2, 16, 8, 19, 5, 13, 10, 17, 14}
	for _, k := range order {
		ix.Insert(k, "v")
	}
	ks := ix.Keys()
	if len(ks) != 20 {
		t.Fatalf("Keys has %d entries, want 20", len(ks))
	}
	for i, k := range ks {
		if k != i+1 {
			t.Fatalf("Keys[%d] = %d, want %d (keys: %v)", i, k, i+1, ks)
		}
	}
}

func TestIndexSplitBoundary(t *testing.T) {
	ix := &Index[int, int]{Max: 4}
	for k := 1; k <= 5; k++ {
		ix.Insert(k, k*10)
	}
	if got := len(ix.Keys()); got != 5 {
		t.Fatalf("Keys has %d entries, want 5", got)
	}
	for k := 1; k <= 5; k++ {
		v, ok := ix.Get(k)
		if !ok || v != k*10 {
			t.Fatalf("Get(%d) = %d, %v, want %d, true", k, v, ok, k*10)
		}
	}
}

func TestIndexScale(t *testing.T) {
	const n = 2000
	start := time.Now()
	ix := &Index[int, int]{Max: 8}
	x := uint32(7)
	seen := make(map[int]bool, n)
	for len(seen) < n {
		x = x*1664525 + 1013904223
		k := int(x>>8) % (4 * n)
		if seen[k] {
			continue
		}
		seen[k] = true
		ix.Insert(k, k)
	}
	ks := ix.Keys()
	if len(ks) != n {
		t.Fatalf("Keys has %d entries, want %d", len(ks), n)
	}
	for i := 1; i < len(ks); i++ {
		if ks[i-1] >= ks[i] {
			t.Fatalf("Keys not strictly ascending at %d: %d then %d", i, ks[i-1], ks[i])
		}
	}
	if d := time.Since(start); d > 3*time.Second {
		t.Fatalf("2k inserts took %v, want under 3s", d)
	}
}
