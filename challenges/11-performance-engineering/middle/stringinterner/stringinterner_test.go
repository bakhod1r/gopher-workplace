package stringinterner

import (
	"strings"
	"testing"
	"unsafe"
)

var sink string

func sameBacking(a, b string) bool {
	return unsafe.StringData(a) == unsafe.StringData(b)
}

func TestInternReturnsTheSameInstance(t *testing.T) {
	var in Interner
	first := in.Intern(strings.Repeat("x", 32))
	second := in.Intern(strings.Repeat("x", 32))
	if first != second {
		t.Fatalf("values differ: %q vs %q", first, second)
	}
	if !sameBacking(first, second) {
		t.Error("interning returned a different backing array; the point is to share one")
	}
}

func TestInternStats(t *testing.T) {
	var in Interner
	in.Intern("a")
	in.Intern("a")
	in.Intern("a")
	in.Intern("b")
	hits, misses := in.Stats()
	if hits != 2 || misses != 2 {
		t.Errorf("Stats = %d, %d, want 2, 2", hits, misses)
	}
	if in.Len() != 2 {
		t.Errorf("Len = %d, want 2", in.Len())
	}
}

func TestInternEmptyString(t *testing.T) {
	var in Interner
	if got := in.Intern(""); got != "" {
		t.Errorf("Intern(\"\") = %q, want empty", got)
	}
	if in.Len() != 1 {
		t.Errorf("Len = %d, want 1", in.Len())
	}
}

func TestInternBytesDoesNotRetainTheCallersBuffer(t *testing.T) {
	var in Interner
	buf := []byte("label")
	got := in.InternBytes(buf)
	copy(buf, "XXXXX")
	if got != "label" {
		t.Errorf("interned string changed to %q when the caller reused their buffer", got)
	}
}

func TestInternBytesHitDoesNotAllocate(t *testing.T) {
	var in Interner
	in.Intern("known-label")
	buf := []byte("known-label")
	allocs := testing.AllocsPerRun(100, func() { sink = in.InternBytes(buf) })
	if allocs != 0 {
		t.Errorf("InternBytes on a known string made %v allocations, want 0", allocs)
	}
}

func TestInternManyRepeats(t *testing.T) {
	var in Interner
	for i := 0; i < 10_000; i++ {
		in.Intern("repeated")
	}
	if in.Len() != 1 {
		t.Errorf("Len = %d, want 1", in.Len())
	}
	hits, misses := in.Stats()
	if hits != 9999 || misses != 1 {
		t.Errorf("Stats = %d, %d, want 9999, 1", hits, misses)
	}
}
