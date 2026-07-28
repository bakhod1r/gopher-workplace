// Package topk returns the k most frequent strings.
package topk

// TopK returns the k most frequent words in xs, most frequent first. Ties are
// broken alphabetically. k is clamped to the number of distinct words.
//
// TODO(candidate): count, then sort by (count desc, word asc), take k.
func TopK(xs []string, k int) []string {
	panic("not implemented")
}
