package batchwriter

import "testing"

func TestFlushesWhenFull(t *testing.T) {
	s := &RecordingSink{}
	w := NewBatchWriter(s, 2)

	w.Write("a")
	if len(s.Batches) != 0 {
		t.Error("flushed before the batch was full")
	}
	w.Write("b")
	if len(s.Batches) != 1 {
		t.Fatalf("Batches = %d, want 1", len(s.Batches))
	}
	if len(s.Batches[0]) != 2 || s.Batches[0][0] != "a" || s.Batches[0][1] != "b" {
		t.Errorf("batch = %v, want [a b]", s.Batches[0])
	}
	if w.Buffered() != 0 {
		t.Errorf("Buffered = %d, want 0", w.Buffered())
	}
}

func TestPartialFlush(t *testing.T) {
	s := &RecordingSink{}
	w := NewBatchWriter(s, 2)

	w.Write("a")
	w.Write("b")
	w.Write("c")
	w.Flush()

	if len(s.Batches) != 2 {
		t.Fatalf("Batches = %d, want 2", len(s.Batches))
	}
	if len(s.Batches[1]) != 1 || s.Batches[1][0] != "c" {
		t.Errorf("second batch = %v, want [c]", s.Batches[1])
	}
}

func TestFlushEmptyIsNoop(t *testing.T) {
	s := &RecordingSink{}
	w := NewBatchWriter(s, 2)
	w.Flush()
	w.Flush()
	if len(s.Batches) != 0 {
		t.Errorf("Batches = %d, want 0", len(s.Batches))
	}
}

func TestNoRecordsLostOrDuplicated(t *testing.T) {
	s := &RecordingSink{}
	w := NewBatchWriter(s, 7)

	const n = 1000
	for i := 0; i < n; i++ {
		w.Write(string(rune('a' + i%26)))
	}
	w.Flush()

	total := 0
	for _, b := range s.Batches {
		total += len(b)
		if len(b) > 7 {
			t.Fatalf("batch of %d exceeds the size limit", len(b))
		}
	}
	if total != n {
		t.Errorf("delivered %d records, want %d", total, n)
	}
}

func TestBufferIsReused(t *testing.T) {
	s := &RecordingSink{}
	w := NewBatchWriter(s, 4)

	for i := 0; i < 4; i++ {
		w.Write("x")
	}
	before := cap(w.buf)
	for i := 0; i < 40; i++ {
		w.Write("y")
	}
	if cap(w.buf) != before {
		t.Errorf("buffer capacity changed from %d to %d; it must be reused", before, cap(w.buf))
	}
}

func TestSizeOneFlushesEveryWrite(t *testing.T) {
	s := &RecordingSink{}
	w := NewBatchWriter(s, 1)
	w.Write("a")
	w.Write("b")
	if len(s.Batches) != 2 {
		t.Errorf("Batches = %d, want 2", len(s.Batches))
	}
}

func BenchmarkWrite(b *testing.B) {
	w := NewBatchWriter(&RecordingSink{}, 64)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		w.Write("record")
	}
}
