// Package batchwriter — Gopher Workplace challenge.
package batchwriter

// Sink receives one batch at a time.
type Sink interface {
	WriteBatch(batch []string)
}

// RecordingSink keeps every batch it received.
type RecordingSink struct {
	Batches [][]string
}

// WriteBatch records the batch.
func (r *RecordingSink) WriteBatch(batch []string) {
	cp := make([]string, len(batch))
	copy(cp, batch)
	r.Batches = append(r.Batches, cp)
}

// BatchWriter buffers records and flushes them in fixed-size batches.
type BatchWriter struct {
	sink Sink
	size int
	buf  []string
}

// NewBatchWriter returns a writer flushing every size records.
func NewBatchWriter(s Sink, size int) *BatchWriter {
	if size < 1 {
		size = 1
	}
	return &BatchWriter{sink: s, size: size, buf: make([]string, 0, size)}
}

// Write buffers a record, flushing when the batch is full.
//
// Examples:
//
//	size 2; Write("a"), Write("b") => one flush of [a b]
func (b *BatchWriter) Write(record string) {
	// TODO(candidate): buffer, then flush when full.
	panic("not implemented")
}

// Flush sends any partial batch. It does nothing when the buffer is empty.
func (b *BatchWriter) Flush() {
	// TODO(candidate): send, then reuse the buffer.
	panic("not implemented")
}

// Buffered returns how many records are waiting.
func (b *BatchWriter) Buffered() int {
	// TODO(candidate): buffered record count.
	panic("not implemented")
}
