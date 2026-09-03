// Package docindexer — Gopher Workplace challenge.
package docindexer

// Doc is one document queued for the search index.
type Doc struct {
	ID   string
	Body string
}

// IndexDocuments indexes every document concurrently and splits the outcome in
// two: the IDs that made it into the index and the IDs that must be replayed.
// Both lists come back sorted. A reindex job is judged by these two lists, so
// neither may depend on which goroutine finished first.
//
// Examples:
//
//	IndexDocuments([]Doc{{"a", "x"}}, index)              => ["a"], []
//	IndexDocuments([]Doc{{"a", "x"}, {"b", ""}}, index)   => ["a"], ["b"]
//	IndexDocuments(nil, index)                            => [], []
func IndexDocuments(docs []Doc, index func(Doc) error) (indexed []string, failed []string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
