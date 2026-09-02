package ringsink

import "testing"

func eq(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestUnderCapacity(t *testing.T) {
	r := NewRingSink(3)
	r.Write("a")
	r.Write("b")
	if r.Len() != 2 {
		t.Errorf("Len = %d, want 2", r.Len())
	}
	if got := r.Snapshot(); !eq(got, []string{"a", "b"}) {
		t.Errorf("Snapshot = %v, want [a b]", got)
	}
}

func TestOverwritesOldest(t *testing.T) {
	r := NewRingSink(3)
	for _, s := range []string{"a", "b", "c", "d"} {
		r.Write(s)
	}
	if r.Len() != 3 {
		t.Errorf("Len = %d, want 3", r.Len())
	}
	if got := r.Snapshot(); !eq(got, []string{"b", "c", "d"}) {
		t.Errorf("Snapshot = %v, want [b c d]", got)
	}
}

func TestExactlyFull(t *testing.T) {
	r := NewRingSink(2)
	r.Write("a")
	r.Write("b")
	if got := r.Snapshot(); !eq(got, []string{"a", "b"}) {
		t.Errorf("Snapshot = %v, want [a b]", got)
	}
}

func TestEmpty(t *testing.T) {
	r := NewRingSink(3)
	if r.Len() != 0 {
		t.Errorf("Len = %d, want 0", r.Len())
	}
	if got := r.Snapshot(); len(got) != 0 {
		t.Errorf("Snapshot = %v, want empty", got)
	}
}

func TestManyWritesStayBounded(t *testing.T) {
	r := NewRingSink(4)
	var s Sink = r
	for i := 0; i < 1_000_000; i++ {
		s.Write("x")
	}
	if r.Len() != 4 {
		t.Errorf("Len = %d, want 4", r.Len())
	}
	if len(r.buf) != 4 {
		t.Errorf("buffer grew to %d, want 4", len(r.buf))
	}
}

func TestWriteDoesNotAllocate(t *testing.T) {
	r := NewRingSink(64)
	line := "a log line"

	avg := testing.AllocsPerRun(1000, func() {
		r.Write(line)
	})
	if avg > 0 {
		t.Errorf("Write allocated %.2f times per call, want 0", avg)
	}
}

func BenchmarkWrite(b *testing.B) {
	r := NewRingSink(1024)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r.Write("line")
	}
}
