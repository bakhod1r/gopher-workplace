// Package ingestlimiter — Gopher Workplace challenge.
package ingestlimiter

// Document is one crawled page waiting to be indexed.
type Document struct {
	ID    string
	Bytes int
}

// IndexDocuments indexes every document concurrently but never lets more than
// maxInFlight index calls run at the same time, because the index host falls
// over above its concurrency budget. A buffered channel of empty structs is
// the semaphore: acquiring a slot is a send, releasing it is a receive.
//
// The result maps document id to the score index returned. maxInFlight below 1
// is clamped to 1.
//
// Examples:
//
//	IndexDocuments([a b c], 2, score) => {a:.., b:.., c:..} with at most 2 concurrent calls
//	IndexDocuments([a], 8, score)     => {a:..}
//	IndexDocuments(nil, 4, score)     => {}
func IndexDocuments(docs []Document, maxInFlight int, index func(Document) int) map[string]int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
