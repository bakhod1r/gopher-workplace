package bloomifc

import (
	"strconv"
	"testing"
)

func TestAddThenContains(t *testing.T) {
	b := NewBloom(1024)
	b.Add("a")
	if !b.MayContain("a") {
		t.Error("MayContain after Add = false")
	}
}

func TestNoFalseNegatives(t *testing.T) {
	b := NewBloom(8192)
	keys := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		k := "key" + strconv.Itoa(i)
		keys = append(keys, k)
		b.Add(k)
	}

	for _, k := range keys {
		if !b.MayContain(k) {
			t.Fatalf("false negative for %q", k)
		}
	}
}

func TestRejectsMostAbsentKeys(t *testing.T) {
	b := NewBloom(65536)
	for i := 0; i < 100; i++ {
		b.Add("present" + strconv.Itoa(i))
	}

	falsePositives := 0
	const n = 1000
	for i := 0; i < n; i++ {
		if b.MayContain("absent" + strconv.Itoa(i)) {
			falsePositives++
		}
	}
	if falsePositives > n/10 {
		t.Errorf("%d false positives out of %d; the filter is not discriminating", falsePositives, n)
	}
}

func TestEmptyFilterRejects(t *testing.T) {
	b := NewBloom(1024)
	if b.MayContain("anything") {
		t.Error("an empty filter should reject")
	}
}

func TestFilterMissing(t *testing.T) {
	b := NewBloom(65536)
	b.Add("a")
	b.Add("b")

	got := FilterMissing(b, []string{"a", "b", "zzz-not-added"})
	if len(got) != 1 || got[0] != "zzz-not-added" {
		t.Errorf("FilterMissing = %v, want [zzz-not-added]", got)
	}
	if n := len(FilterMissing(b, nil)); n != 0 {
		t.Errorf("FilterMissing(nil) len = %d, want 0", n)
	}
}

func TestMemoryIsFixed(t *testing.T) {
	b := NewBloom(4096)
	before := len(b.bits)
	for i := 0; i < 100000; i++ {
		b.Add("k" + strconv.Itoa(i))
	}
	if len(b.bits) != before {
		t.Errorf("bitset grew from %d to %d bytes", before, len(b.bits))
	}
}

func BenchmarkAdd(b *testing.B) {
	f := NewBloom(65536)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f.Add("key")
	}
}
